# AIRSPACE DEMO API
### TIM LIEBERMAN   : )

#### OBJECTIVE: 
Using N_S_UAS_Flight_Restriction data create a api that takes a lat long and responds with message revealing if cords
are in restricted airspace or not. 

### PRE-Reqs
#### Required To Run:
I'm presuming you already have these tools installed but if not i've provided some psuedo installs.  They might be accurate.

- `brew install docker`
- `brew install go`

#### Testing: 
The only prerequisite for testing is that you have gomock installed... 
- ```go install github.com/golang/mock/mockgen@latest```

### Run Local Build

1. From project root:
   2. `docker compose up --build`
      3. No Errors ? http://localhost:8080/restricted-airspace/32.3372/-84.9914 : Try again
     
         4. **Note**:  expect these.. _airspace-challenge   | 2024/11/16 22:23:55 /app/internal/database/db.go:53
            airspace-challenge   | [error] failed to initialize database, got error failed to connect to `host=postgres user=postgres database=postgres`: dial error (dial tcp 172.20.0.2:5432: connect: connection refused)
            airspace-challenge   | 2024/11/16 22:23:55 failed to connect to `host=postgres user=postgres database=postgres`: dial error (dial tcp 172.20.0.2:5432: connect: connection refused)_

### Test Local
1. Form root bash:
   2. `go test ./...`


### How it works:  
The API was built using a combination of Go, Docker, and PostgreSQL to provide a scalable and maintainable solution for 
handling flight restriction data.  The use of the Echo web framework allows for a straightforward implementation of 
RESTful endpoints, and it provides built-in tools for routing and middleware management. The handler method uses the
Echo framework to process the coordinates, validate them, and then query the mock repository for the restricted airspace
status. This approach is lightweight and effective for handling real-time data requests.

Docker was employed to containerize the services, ensuring consistency across different environments. The Docker 
Compose configuration defines the application stack, which consists of three primary services: the airspace-challenge API 
service, the data-loader service for populating the database with flight data, and the postgres service for storing data.
Docker simplifies the deployment process by encapsulating the application and its dependencies into isolated containers,

The API also relies on the PostGIS extension for PostgreSQL, which enables spatial data processing and geographic 
queries. PostGIS is a powerful spatial database extension that allows PostgreSQL to handle geographic objects, such as
points, polygons, and lines, which is essential for this API's functionality, specifically when working with flight 
restriction zones.

The core feature utilized here is the ability to store and query spatial data using PostGIS functions like 
ST_Intersects. This function allows the API to determine if a specific point (representing a location on the Earth, 
defined by latitude and longitude) intersects with any defined restricted airspace area stored in the database. 
By leveraging PostGIS, I'm able to efficiently perform geospatial queries, such as checking whether a given set of 
coordinates lies within a restricted airspace zone.

### What features would I add to this.

1. **Geospatial Search Enhancements**
   Distance Calculations: Include functionality to calculate the distance between a coordinate and the nearest restricted airspace zone. This could be helpful for proximity-based queries (e.g., "Is this location close to any restricted airspace?").
2. **Buffer Zones:** Allow for the creation of buffer zones around restricted airspace, where users can query if a point falls within a certain distance of a restricted area.
3. **Real-Time Alerts:** Add a system for notifying users when a specific area becomes restricted or when an existing restriction is lifted. Notifications could be pushed to users via email or SMS.
4. **Geospatial Visualization:** Allow users to visualize flight restrictions on a map via a web interface. This could be useful for external systems that need to display real-time information about restricted airspace on a map.
5. **API Keys:** Implement API key-based authentication to ensure that only authorized users can access certain parts of the API.


### Full log of successful build... 

2024-11-16 17:23:55 airspace-challenge | 2024-11-16 17:23:55 airspace-challenge | 2024/11/16 22:23:55 /app/internal/database/db.go:53 2024-11-16 17:23:55 airspace-challenge | [error] failed to initialize database, got error failed to connect to host=postgres user=postgres database=postgres: dial error (dial tcp 172.20.0.2:5432: connect: connection refused) 2024-11-16 17:23:55 airspace-challenge | 2024/11/16 22:23:55 failed to connect to host=postgres user=postgres database=postgres: dial error (dial tcp 172.20.0.2:5432: connect: connection refused) 2024-11-16 17:23:55 airspace-challenge | 2024-11-16 17:23:55 airspace-challenge | 2024/11/16 22:23:55 /app/internal/database/db.go:53 2024-11-16 17:23:55 airspace-challenge | [error] failed to initialize database, got error failed to connect to host=postgres user=postgres database=postgres: dial error (dial tcp 172.20.0.2:5432: connect: connection refused) 2024-11-16 17:23:55 airspace-challenge | 2024/11/16 22:23:55 failed to connect to host=postgres user=postgres database=postgres: dial error (dial tcp 172.20.0.2:5432: connect: connection refused) 2024-11-16 17:23:55 airspace-challenge | 2024-11-16 17:23:55 airspace-challenge | 2024/11/16 22:23:55 /app/internal/database/db.go:53 2024-11-16 17:23:55 airspace-challenge | [error] failed to initialize database, got error failed to connect to host=postgres user=postgres database=postgres: dial error (dial tcp 172.20.0.2:5432: connect: connection refused) 2024-11-16 17:23:55 airspace-challenge | 2024/11/16 22:23:55 failed to connect to host=postgres user=postgres database=postgres: dial error (dial tcp 172.20.0.2:5432: connect: connection refused) 2024-11-16 17:23:56 airspace-challenge | 2024-11-16 17:23:56 airspace-challenge | 2024/11/16 22:23:56 /app/internal/database/db.go:53 2024-11-16 17:23:56 airspace-challenge | [error] failed to initialize database, got error failed to connect to host=postgres user=postgres database=postgres: dial error (dial tcp 172.20.0.2:5432: connect: connection refused) 2024-11-16 17:23:56 airspace-challenge | 2024/11/16 22:23:56 failed to connect to host=postgres user=postgres database=postgres: dial error (dial tcp 172.20.0.2:5432: connect: connection refused) 2024-11-16 17:23:57 airspace-challenge | 2024-11-16 17:23:57 airspace-challenge | 2024/11/16 22:23:57 /app/internal/database/db.go:53 2024-11-16 17:23:57 airspace-challenge | [error] failed to initialize database, got error failed to connect to host=postgres user=postgres database=postgres: dial error (dial tcp 172.20.0.2:5432: connect: connection refused) 2024-11-16 17:23:57 airspace-challenge | 2024/11/16 22:23:57 failed to connect to host=postgres user=postgres database=postgres: dial error (dial tcp 172.20.0.2:5432: connect: connection refused) 2024-11-16 17:23:55 data-loader | Waiting for postgres:5432 to be available... 2024-11-16 17:23:55 postgres-1 | The files belonging to this database system will be owned by user "postgres". 2024-11-16 17:23:59 data-loader | Connection to postgres (172.20.0.2) 5432 port [tcp/postgresql] succeeded! 2024-11-16 17:23:59 data-loader | postgres:5432 is available. Starting the command... 2024-11-16 17:23:59 data-loader | 2024/11/16 22:23:59 User: postgres 2024-11-16 17:24:00 data-loader | Data successfully loaded. 2024-11-16 17:23:58 airspace-challenge | Server running on port 8080 2024-11-16 17:23:58 airspace-challenge | 2024-11-16 17:23:55 postgres-1 | This user must also own the server process. 2024-11-16 17:23:55 postgres-1 | 2024-11-16 17:23:55 postgres-1 | The database cluster will be initialized with locale "en_US.utf8". 2024-11-16 17:23:55 postgres-1 | The default database encoding has accordingly been set to "UTF8". 2024-11-16 17:23:55 postgres-1 | The default text search configuration will be set to "english". 2024-11-16 17:23:55 postgres-1 | 2024-11-16 17:23:55 postgres-1 | Data page checksums are disabled. 2024-11-16 17:23:55 postgres-1 | 2024-11-16 17:23:58 airspace-challenge | ____ __ 2024-11-16 17:23:58 airspace-challenge | / // / ___ 2024-11-16 17:23:58 airspace-challenge | / // / _ / _
2024-11-16 17:23:58 airspace-challenge | //_////___/ v4.12.0 2024-11-16 17:23:58 airspace-challenge | High performance, minimalist Go web framework 2024-11-16 17:23:58 airspace-challenge | https://echo.labstack.com 2024-11-16 17:23:58 airspace-challenge | _____________________________O/ 2024-11-16 17:23:58 airspace-challenge | O
2024-11-16 17:23:58 airspace-challenge | ⇨ http server started on [::]:8080 2024-11-16 17:24:13 airspace-challenge | 2024/11/16 22:24:13 slat : before parse 32.3372, 32.3372 2024-11-16 17:24:13 airspace-challenge | {"time":"2024-11-16T22:24:13.922775222Z","id":"","remote_ip":"192.168.65.1","host":"localhost:8080","method":"GET","uri":"/restricted-airspace/32.3372/-84.9914","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36","status":200,"error":"","latency":58126833,"latency_human":"58.126833ms","bytes_in":0,"bytes_out":78} 2024-11-16 17:24:14 airspace-challenge | {"time":"2024-11-16T22:24:14.045538125Z","id":"","remote_ip":"192.168.65.1","host":"localhost:8080","method":"GET","uri":"/favicon.ico","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36","status":404,"error":"code=404, message=Not Found","latency":18833,"latency_human":"18.833µs","bytes_in":0,"bytes_out":24} 2024-11-16 17:24:59 airspace-challenge | 2024/11/16 22:24:59 slat : before parse 32.3372, 32.3372 2024-11-16 17:24:59 airspace-challenge | {"time":"2024-11-16T22:24:59.03224459Z","id":"","remote_ip":"192.168.65.1","host":"localhost:8080","method":"GET","uri":"/restricted-airspace/32.3372/-84.9914","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36","status":200,"error":"","latency":18202667,"latency_human":"18.202667ms","bytes_in":0,"bytes_out":78} 2024-11-16 17:24:59 airspace-challenge | {"time":"2024-11-16T22:24:59.092838757Z","id":"","remote_ip":"192.168.65.1","host":"localhost:8080","method":"GET","uri":"/restricted-airspace/32.3372/-84.9914","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36","status":200,"error":"","latency":2040709,"latency_human":"2.040709ms","bytes_in":0,"bytes_out":78} 2024-11-16 17:24:59 airspace-challenge | {"time":"2024-11-16T22:24:59.361158423Z","id":"","remote_ip":"192.168.65.1","host":"localhost:8080","method":"GET","uri":"/favicon.ico","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36","status":404,"error":"code=404, message=Not Found","latency":30542,"latency_human":"30.542µs","bytes_in":0,"bytes_out":24} 2024-11-16 17:24:59 airspace-challenge | {"time":"2024-11-16T22:24:59.411124382Z","id":"","remote_ip":"192.168.65.1","host":"localhost:8080","method":"GET","uri":"/favicon.ico","user_agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36","status":404,"error":"code=404, message=Not Found","latency":11916,"latency_human":"11.916µs","bytes_in":0,"bytes_out":24}
