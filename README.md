<h1 align="center">
  <a href="https://github.com/r7wx/easy-gate"><img width="90" src="assets/logo.svg" /></a>
  <br />
  easy-gate
</h1>
<h4 align="center">A gate to your self hosted infrastructure</h3>

<p align="center">
<a href="https://github.com/r7wx/easy-gate/actions/workflows/build.yml">
<img src="https://github.com/r7wx/easy-gate/actions/workflows/build.yml/badge.svg"/></a>
<a href="https://github.com/r7wx/easy-gate/actions/workflows/test.yml">
<img src="https://github.com/r7wx/easy-gate/actions/workflows/test.yml/badge.svg"/></a>
<a href="https://www.codefactor.io/repository/github/r7wx/easy-gate"><img src="https://www.codefactor.io/repository/github/r7wx/easy-gate/badge?s=3531cadc71a73618609de429289d9524a38c33cd" alt="CodeFactor" /></a>
</p>

---

<img src="assets/screenshot.png" />

<p align="justify">
Easy Gate is a simple web application built in Go and React that acts as the home page for your self-hosted infrastructure. Services and notes are parsed from a JSON file in real-time (without restarting the application). Services and notes can be assigned to one or more groups to show items only to specific users (based on their IP addresses).
</p>

## Deployment

<p align="justify">
Easy Gate can be deployed on any server that supports Docker. If your infrastructure is running behind a configured nginx instance, it is recommended to read the Docker Compose (with nginx) section.
</p>

### Docker

<p align="justify">
You can deploy an instance of easy-gate by using docker:
</p>

```bash
docker run -d --name=easy-gate \
  -p 8080:8080 \
  -v /path/to/easy-gate.json:/etc/easy-gate/easy-gate.json \
  --restart unless-stopped \
  r7wx/easy-gate
```

### Docker Compose

You can also run easy-gate by using the provided docker-compose file:

```bash
docker-compose up
```

### Docker Compose (with nginx)

<p align="justify">
If you need to host easy-gate behind an already installed nginx instance, you can use the docker-compose file in the examples directory:
</p>

```bash
easy-gate:
    image: r7wx/easy-gate:latest
    container_name: easy-gate
    expose:
      - 8080
    networks:
      - nginx-net
    volumes:
      - ../easy-gate.json:/etc/easy-gate/easy-gate.json
      - ./easy-gate.nginx.conf:/etc/nginx/conf.d/default.conf

[...]
```

<p align="justify">
In order to correctly use the groups feature you must overwrite the internal easy-gate nginx configuration in order to forward requests headers:
</p>

```bash
[...]
location /api {
  proxy_redirect off;
  proxy_pass_request_headers on;
  proxy_pass http://127.0.0.1:8081;
}
[...]
```

You can find the complete file in the examples directory (examples/easy-gate.nginx.conf).

## Configuration

<p align="justify">
Easy gate can be configured by the easy-gate.json file. An illustrative configuration file is provided in the root directory of this repository (easy-gate.json).
</p>

### Groups

<p align="justify">
Group entries are used to define which users can see which items, by providing the user subnet:
</p>

```json
"groups": [
  {
    "name": "internal",
    "subnet": "192.168.1.1/24"
  },
  {
    "name": "vpn",
    "subnet": "10.8.1.1/24"
  }
]
```

### Services

<p align="justify">
A service entry is used to define a service that is available in the infrastructure. Each service has a name, an url, an icon and the groups that can see it (defined in the groups section). If no group is provided the item can be seen by all users:
</p>

```json
{
  "icon": "fa-brands fa-git-square",
  "name": "Git",
  "url": "https://git.example.vpn",
  "groups": [
      "vpn"
  ]
},
{
  "icon": "fa-brands fa-docker",
  "name": "Portainer",
  "url": "https://portainer.example.internal",
  "groups": []
}
```

### Notes

<p align="justify">
A note entry is used to define a simple text note which has a title and a content. Each note has a name, the note content (text) and the groups that can see it (defined in the groups section). If no group is provided the item can be seen by all users:
</p>

```json
{
  "name": "Simple note",
  "text": "This is a simple note for vpn users",
  "groups": [
    "vpn"
  ]
},
{
  "name": "Global note",
  "text": "This note will be visible to everyone",
  "groups": []
}
```

### Icons

Icons are provided by the [Font Awesome](https://fontawesome.com/icons?d=gallery) library. Get the appropriate icon name by using the Font Awesome website (only free icons are available).
