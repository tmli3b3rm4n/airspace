# AIRSPACE DEMO API

#### OBJECTIVE: 
Go Engineering Challenge:
- Download the following dataset as GeoJSON - https://udds-faa.opendata.arcgis.com/datasets/faa::national-security-uas-flight-restrictions-1/explore - This is the offical FAA dataset for the National Security UAS Flight Restrictions.  It contains collection of polygons representing the boundaries of areas that a drone operator cannot fly a drone.
- Create an API endpoint that accepts a latitude and longitude and returns a positive or negative response indicating whether a drone operator can fly a drone at that location.
- You must right the API code using Go, but you are free to decide how to model that API: Rest or gRPC is an acceptable solution.
- Provide documentation on how to run the application.

Don't spend too much time optimizing your code, but think about what possible enhancements you would make to it.
Submit your code via any means to kelly.bigley@airspacelink.com .  If you decide to prepare a repository online (ie: Github), please do NOT make it public.  Rather keep the repository private and send an invitation as a collaborator.

#### Solution
Using the National_Security_UAS_Flight_Restrictions.geojson data, I created an API that takes latitude and longitude as inputs and responds with a message indicating whether the coordinates fall within restricted airspace. While not strictly necessary, I also built a React application with a map feature that displays the restricted airspace using the provided geometry. The React app consumes the Go API and allows users to click on the map to select coordinates, which can then be checked against the API. Although I wouldn't normally include the frontend in such a project, I wanted to make it as straightforward as possible for you to run.

#### Prerequisites
I'm presuming you already have these tools installed but if not i've provided installs.

- `brew install docker`
- `brew install go`
-
#### Installing
- `mkdir -pv ~/go/src/github.com/tmli3b3rm4n/ && cd ~/go/src/github.com/tmli3b3rm4n && git clone git@github.com:tmli3b3rm4n/airspace.git && cd airspace`

#### Testing: 
The only prerequisite for testing is that you have gomock installed... 
- `go install github.com/golang/mock/mockgen@latest`

#### Swagger Docs
 * `go install github.com/swaggo/swag/cmd/swag@latest`
 * `swag init -g cmd/airspace_challenge/main.go`

#### Run Local Build
* `docker compose up --build`
*  sanity check  http://localhost:8080/restricted-airspace/32.3372/-84.9914
* http://localhost:8080/swagger//index.html#/

#### Frontend
* http://localhost:3005/

#### Tests
Tests are designed to run in any environment—local, cloud, or otherwise. I achieve this by using interfaces to mock the database, which abstracts the application logic from the actual database implementation. 

* `go test ./...`


#### How it works:  
The API is built using Go, Docker, and PostgreSQL to manage flight restriction data. Echo framework handles RESTful endpoints, providing routing and middleware tools. The handler method processes coordinates, validates them, and queries the database for restricted airspace status. For testing, mock data is used to test handlers locally.

Docker is used to containerize services, ensuring environment consistency. The stack includes the API service, a data-loader for populating the database, and a PostgreSQL service. Docker Compose handles deployment across containers.

PostGIS, an extension for PostgreSQL, enables spatial data processing and geographic queries. It supports geographic objects like points, polygons, and lines, critical for working with flight restriction zones. The ST_Intersects function is used to check if a point intersects with any restricted airspace zone, allowing efficient geospatial queries.

### What features would I add to this.

1. **Geospatial Search Enhancements**
   Distance Calculations: Include functionality to calculate the distance between a coordinate and the nearest restricted airspace zone. This could be helpful for proximity-based queries (e.g., "Is this location close to any restricted airspace?").
2. **Buffer Zones:** Allow for the creation of buffer zones around restricted airspace, where users can query if a point falls within a certain distance of a restricted area.
3. **Real-Time Alerts:** Add a system for notifying users when a specific area becomes restricted or when an existing restriction is lifted. Notifications could be pushed to users via email or SMS.
4. **Geospatial Visualization:** Allow users to visualize flight restrictions on a map via a web interface. This could be useful for external systems that need to display real-time information about restricted airspace on a map.
5. **API Keys:** Implement API key-based authentication to ensure that only authorized users can access certain parts of the API.
6. **Github Actions**
