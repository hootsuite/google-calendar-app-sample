//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=cfg.yaml -include-tags=PlannedContent ../../tools/openapi.yaml
package planned_content

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strings"

	"golang.org/x/oauth2"
	cal "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const (
	YYYYMMDD         = "2006-01-02"
	YYYYMMDDTHHMMSS  = "2006-01-02T15:04:05"
	YYYYMMDDTHHMMSSZ = "2006-01-02T15:04:05Z"
)

type Server struct {
	config *oauth2.Config
}

func NewServer(config *oauth2.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) GetPlannedContent(w http.ResponseWriter, r *http.Request, params GetPlannedContentParams) {
	ctx := r.Context()
	// Extract and log the Bearer token
	at := strings.Replace(params.Authorization, "Bearer ", "", 1)
	t := &oauth2.Token{
		AccessToken: at,
		TokenType:   "Bearer",
	}

	// Create the calendar service
	calService, err := cal.NewService(ctx, option.WithTokenSource(s.config.TokenSource(ctx, t)))
	if err != nil {
		log.Printf("Error creating calendar service: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Respond with planned content
	var resp PlannedContentResponse = PlannedContentResponse{}

	if params.ParentId == nil || len(*params.ParentId) == 0 {
		list, err := calService.CalendarList.List().Do()
		if err != nil {
			log.Printf("Error listing calendars: %v\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Process calendars
		resp.Data = ConvertCalendars(list)
	} else {
		// List events
		getEvents := calService.Events.List(*params.ParentId).EventTypes("default").MaxResults(5)

		if params.Cursor != nil {
			getEvents = getEvents.PageToken(*params.Cursor)
		}

		if params.StartDate != nil {
			sd := params.StartDate.Format(YYYYMMDDTHHMMSSZ)
			getEvents = getEvents.TimeMin(sd)
		}
		if params.EndDate != nil {
			ed := params.EndDate.Format(YYYYMMDDTHHMMSSZ)
			getEvents = getEvents.TimeMax(ed)
		}

		events, err := getEvents.Do()
		if err != nil {
			log.Printf("Error listing events: %v\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Process events
		resp.Data = ConvertEvents(events)
		resp.Meta = &Meta{
			Cursor: &events.NextPageToken,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding response: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ConvertEvents(events *cal.Events) []PlannedContentItem {
	content := PlannedContent

	// Process events
	result := make([]PlannedContentItem, 0, len(events.Items))

	for _, i := range events.Items {
		contentItem := PlannedContentItem{
			Id:        i.Id,
			Name:      &i.Summary,
			MediaType: &content,
			Permalink: &i.HtmlLink,
			Body:      &i.Description,
			Dates: &Dates{
				Due:     &i.End.DateTime,
				Created: &i.Created,
				Updated: &i.Updated,
			},
		}
		if i.Creator != nil {
			author := AuthorMetadata{
				Id: &i.Creator.Id,
			}
			if len(i.Creator.DisplayName) > 0 {
				if strings.ContainsAny(i.Creator.DisplayName, " ") {
					author.FirstName = &(strings.Split(i.Creator.DisplayName, " ")[0])
					author.LastName = &(strings.Split(i.Creator.DisplayName, " ")[1])
				} else {
					author.FirstName = &i.Creator.DisplayName
				}
			}
			contentItem.Author = &author
		}
		if len(i.Attachments) > 0 {
			var attachments []Attachment
			for _, a := range i.Attachments {
				if len(a.MimeType) > 0 {
					if mimeType, err := IsMimeTypeAllowed(a.MimeType); err != nil {
						log.Printf("Error processing attachment: %v\n", err)
					} else {
						attachment := Attachment{
							Id:        a.FileId,
							Name:      &a.Title,
							MimeType:  &mimeType,
							MediaType: GetMediaType(mimeType),
							Original: &ImageMetadata{
								Url: &a.FileUrl,
							},
							Thumbnail: &ThumbnailMetadata{
								Url: &a.IconLink,
							},
						}

						attachments = append(attachments, attachment)
					}
				}

			}
			contentItem.Attachments = &attachments
		}
		result = append(result, contentItem)
	}
	return result
}

func ConvertCalendars(list *cal.CalendarList) []PlannedContentItem {
	folder := Folder
	result := make([]PlannedContentItem, 0, len(list.Items))
	for _, c := range list.Items {
		item := PlannedContentItem{
			Id:        c.Id,
			Name:      &c.Summary,
			MediaType: &folder,
		}
		result = append(result, item)
	}
	return result
}

// DetermineMimeType determines the mime type of an image based on the file extension
func IsMimeTypeAllowed(mimeType string) (string, error) {
	allowedMimeTypes := []string{
		"image/jpeg",
		"image/png",
		"image/gif",
		"image/webp",
		"video/mp4",
		"video/quicktime",
		"application/pdf",
	}

	if !slices.Contains(allowedMimeTypes, mimeType) {
		log.Println("Invalid mime type: ", mimeType)
	}

	return mimeType, nil
}

func GetMediaType(mimeType string) *string {
	var (
		image = "image"
		video = "video"
		pdf   = "pdf"
	)
	if strings.Contains(mimeType, "image") {
		return &image
	} else if strings.Contains(mimeType, "video") {
		return &video
	} else if strings.Contains(mimeType, "application/pdf") {
		return &pdf
	}
	return nil
}
