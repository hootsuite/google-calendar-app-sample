# Google Calendar Sample App
- [Google Calendar Sample App](#google-calendar-sample-app)
  - [Requirements](#requirements)
  - [Set-up Google](#set-up-google)
    - [Set-up an new Google Project](#set-up-an-new-google-project)
    - [Get Calendar API Access](#get-calendar-api-access)
    - [Get access to credentials](#get-access-to-credentials)
  - [Set-up Hootsuite](#set-up-hootsuite)
  - [Make Commands](#make-commands)
  - [Host](#host)


## Requirements
- This project uses Go Lang version 1.22. You can view the instructions on how to install it [here](https://go.dev/doc/install).
- Access to a Google Account
- Uses the latest OpenAPI doc which you can find [here](https://app-directory.s3.amazonaws.com/docs/outbound-api/redoc-static.html)

## Set-up Google
We will be connecting to the google calendar you will need to have and account there and set-up access to Google Calendar.  You can read up about the [Google Calendar API](https://developers.google.com/calendar).

### Set-up an new Google Project
To create a new API project go to [Google Console](https://console.cloud.google.com/)

1. Beside the Google Cloud icon in the top right click on the drop down and click on `New Project`.
2. Name the project `google-calendar-sample`. You do not have to connect it to an organisation and click next.

We can head to the project api dashboard found [here](hhttps://console.cloud.google.com/apis/dashboard).  Make sure the dropdown is still selecting your project.

### Get Calendar API Access
From the [API dashboard](https://console.cloud.google.com/apis/dashboard).
1. Click on `Library` 
2. Click on `Google Calendar API` 
3. Click on `Enable`.  This will turn on the calendar APIs

### Get access to credentials
This will allow people to log into their accounts to access their own google calendar
From the [API dashboard](https://console.cloud.google.com/apis/dashboard).
1. Click on `Credentails`
2. Click on `+ Create Credentails` and `OAuth Client ID`
3. Click on `Configure Consent Screen`
   1. Click on `External` this will allow people to use the project
   2. Fill in the form with
      - App Name: `Hootsuite Sample Calendar App`
      - User Support Email: `<Your email>`
      - Authorised Domains: `hootsuite.com`
      - Contact Information: `<Your email>`
   3. Click `Save`
   4. Click on `Add or Remove Scopes` and add the following:
      -  `../auth/calendar.readonly`
      -  Click on `Update`
4. Click on `Save and Continue`
5. Now add Test Users that you want to test out your app (probably want to add yourself to this list)
6. Click on `Save and Continue`
7. Click on `+ Create Credentails` and `OAuth Client ID` again.
   1. Application Type: `Web application`
      - Name: `Sample App`
   2. We will have to come back to this once we have more information for the redirect url
   3. Click `Save`
   4. Copy the `Client ID` and `Client Secret` we will need these to set-up on the Hootsuite side

## Set-up Hootsuite
Once your developer account is active you can go to [My Apps](https://hootsuite.com/developers/my-apps)
1. Click on `Create New App`
2. Fill out the form:
   - App Title: `Google Calendar Sample App`
   - Description: `This is a sample app to connect to Planner Content to Google Calendar`
   - The rest can be black for the sample
3. Click on the app you just created
4. Click on `Edit`
5. In the `Extension Authentication (OAuth 2.0)` section.
   - Client ID: This is the id from Google `<pattern>.apps.googleusercontent.com`
   - Client Secret: This is the secret from Google `GOCSPX-<pattern>`
   - Token URL: From Google `https://oauth2.googleapis.com/token`
   - Authorization URL: From Google `https://accounts.google.com/o/oauth2/auth`
   - Scope: From Google `https://www.googleapis.com/auth/calendar.readonly`
6. Click `Save`
7. Copy `Redirect URI` that was generated will need to add this to the Google Settings
   1. Go to the [Credentails](https://console.cloud.google.com/apis/credentials) in Google
   2. Click on `Sample App` under `Oauth 2.0 Client IDs`
   3. Click `+ Add URI` under `Authorised redirect URIs`
   4. Paste the Redirect URI from Hootsuite here
8. Click on `New App Extension` 
    - App Extension Type: `Planner`
    - App Extension Title: `Sample Google Calendar App`
9.  CLick `Add`

## Make Commands
There is 3 commands in the `Makefile`:
 - `generate`: Will re-generate the `.gen.go` files
 - `build`: Will compile the code into a binary
 - `run`: will run the code locally

## Host

You can see how to set-up the app [Heroku](http://www.heroku.com/) with this [guide](https://developer.hootsuite.com/docs/iframe-sdk-sample-apps)
