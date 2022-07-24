# FamTube

FamTube is a service which keeps updated youtube videos of a predefined query and provides the videos of search filter.
The videos are sorted in descending order of published datetime so you will never miss the latest video on your favourite subject.

---

## Prerequisites:
System should have docker and docker-compose installed. 

References:
  - https://docs.docker.com/engine/install/
  - https://docs.docker.com/compose/install/

## Getting Started

1. Clone this repo: `git clone https://github.com/99neerajsharma/FamTube`
2. Go to the FamTube directory using `cd FamTube`
3. Edit the youtube api key in config.yml for `yt_keys` under `worker` and edit the `query` value for your search query
4. Once above change are done use `docker-compose up --build` command to run the services.
5. Checkout the server status at `localhost:3000` in the browser
6. Use the `FamTube.postman_collection.json` in postman for API details.

---

