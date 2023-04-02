# YarnSwap - Final Project Higher Diploma in Computer Science 2023
# Project Name: Yarn Swap
## Project Description: 
Yarn Swap is a Progressive Web Application where crafters can list and view yarns and can swap or procure yarns from others that could otherwise be forgotten or wasted. This project is the backend API for the Yarn Swap application.

### The following technologies have been used in the making of the Yarn Swap backend API:
* Golang
* Gin Web Framework
* Firebase (Auth, Realtime Database, Storage)

## Project Status: Deployed
* Red Hat OpenShift Cluster: http://yarnswap-yarn-swap.apps.fwaters.uw4y.s1.devshift.org/
**This may not be currently running.
* The corresponding frontend application is deployed here: http://yarnswap-fe-yarn-swap.apps.fwaters.uw4y.s1.devshift.org/

## Features
* There are 11 API endpoints adding and retrieving data on Firebase Realtime Database.

## Setup Requirements
* Clone this repo
* Open it
* Set up a firebase account and add required credentials.
* From the root directory run `go mod download && go mod verify`
* From the root directory run `go run .`

## Testing
* API endpoint testing has been added using the Go Testing Package.
* Run these tests with `go test`

## Final Project Report
* For more information please refer to the final project report
  - https://github.com/Fiona-Waters/YarnSwap/blob/main/FionaWaters-20095357-Project-FinalReport.pdf
