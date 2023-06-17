<h1 align="center">
  <a href="https://github.com/r7wx/easy-gate"><img width="90" src="assets/logo.svg" /></a>
  <br />
  easy-gate
</h1>
<h4 align="center">A gate to your self hosted infrastructure</h3>

<p align="center">
<img src="https://img.shields.io/github/v/release/r7wx/easy-gate" alt="Release" />
<a href="https://github.com/r7wx/easy-gate/actions/workflows/build.yml">
<img src="https://github.com/r7wx/easy-gate/actions/workflows/build.yml/badge.svg"/></a>
<a href="https://github.com/r7wx/easy-gate/actions/workflows/test.yml">
<img src="https://github.com/r7wx/easy-gate/actions/workflows/test.yml/badge.svg"/></a>
<a href="https://www.codefactor.io/repository/github/r7wx/easy-gate"><img src="https://www.codefactor.io/repository/github/r7wx/easy-gate/badge?s=3531cadc71a73618609de429289d9524a38c33cd" alt="CodeFactor" /></a>
</p>

---

<img src="assets/demo.png" />

<p align="justify">
Easy Gate is a simple web application designed to serve as the central hub for your self-hosted infrastructure. With real-time parsing of services and notes from a configuration file, Easy Gate ensures seamless updates without the need to restart the application. Moreover, it offers the flexibility to assign items to specific user groups, allowing tailored access based on their IP addresses.
</p>

### Features

- Real-time parsing of services and notes from a configuration file (JSON/YAML) without requiring application restarts.
- Flexible assignment of services and notes to one or more user groups, ensuring items are displayed exclusively to specific users based on their IP addresses.
- Organization of services into categories.
- Customizable theme and icons.
- Versatile deployment options with the choice to run Easy Gate as a standalone executable, free from external dependencies, or as a Docker container.

## Deployment

### Standalone Executable

<p align="justify">
In order to run Easy Gate as a standalone executable, you can build it from source code or download a pre-built binary from the latest release.
</p>

**Build from source:**

```bash
git clone https://github.com/r7wx/easy-gate.git
cd easy-gate
make
```

**Run executable:**

```bash
easy-gate <path-to-config-file>
```

<p align="justify">
Configuration file can be either a JSON or a YAML file.
</p>

### Docker

<p align="justify">
You can deploy an instance of Easy Gate by using Docker:
</p>

```bash
docker run -d --name=easy-gate \
  -p 8080:8080 \
  -v /path/to/easy-gate.json:/etc/easy-gate/easy-gate.json \
  --restart unless-stopped \
  r7wx/easy-gate:latest
```

<p align="justify">
By default the Easy Gate image will look for a configuration file in /etc/easy-gate/easy-gate.json, but this value can be overridden by using the EASY_GATE_CONFIG_PATH environment variable:
</p>

```bash
docker run -d --name=easy-gate \
  -p 8080:8080 \
  -v /path/to/easy-gate.yml:/another/path/easy-gate.yml \
  --restart unless-stopped \
  -e EASY_GATE_CONFIG_PATH=/another/path/easy-gate.yml \
  r7wx/easy-gate:latest
```

### Docker Compose

You can run Easy Gate by using the provided docker-compose file:

```yml
services:
  easy-gate:
    image: r7wx/easy-gate:latest
    build: .
    container_name: easy-gate
    restart: unless-stopped
    ports:
      - 8080:8080
    volumes:
      - ./easy-gate.json:/etc/easy-gate/easy-gate.json
```

```bash
docker-compose up
```

<p align="justify">
By default the Easy Gate image will look for a configuration file in /etc/easy-gate/easy-gate.json, but this value can be overridden by using the EASY_GATE_CONFIG_PATH environment variable:
</p>

```yml
services:
  easy-gate:
    image: r7wx/easy-gate:latest
    build: .
    container_name: easy-gate
    restart: unless-stopped
    environment:
      - EASY_GATE_CONFIG_PATH=/etc/easy-gate/easy-gate.yml
    ports:
      - 8080:8080
    volumes:
      - ./easy-gate.yml:/etc/easy-gate/easy-gate.yml
```

### Docker Compose (behind proxy)

<p align="justify">
If you need to host Easy Gate behind an already running nginx instance (or other reverse proxies), you can use the docker-compose file in the examples directory:
</p>

```yml
services:
  easy-gate:
    image: r7wx/easy-gate:latest
    container_name: easy-gate
    expose:
      - 8080
    networks:
      - nginx-net
    volumes:
      - ../easy-gate.json:/etc/easy-gate/easy-gate.json

  nginx:
    image: nginx:latest
    container_name: nginx
    ports:
      - 80:80
    networks:
      - nginx-net
    volumes:
      - ./nginx.conf:/etc/nginx/conf.d/default.conf

networks:
  nginx-net:
    driver: bridge
```

<p align="justify">
In order to correctly use the groups feature, the nginx instance (or your other reverse proxy) must be configured to use the X-Forwarded-For header:
</p>

```nginx
[...]
location / {
  proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
  proxy_pass http://easy-gate:8080;
}
[...]
```

It is also mandatory to set "behind_proxy" to true in your easy-gate configuration file:

```json
[...]
"behind_proxy": true,
[...]
```

You can find the complete docker-compose and nginx configuration files in the examples directory. The same logic applies to standalone and Docker deployments.

## Configuration

<p align="justify">
Easy gate can be configured by a JSON or a YAML configuration file. An example configuration is provided in the root directory of this repository (easy-gate.json/easy-gate.yml).
</p>

### Options

- **addr:** IP address to listen on
- **use_tls:** If true, the application will use TLS
- **cert_file:** Path to the SSL certificate file (if TLS is enabled)
- **key_file:** Path to the SSL key file (if TLS is enabled)
- **behind_proxy:** If true, the application will use the X-Forwarded-For header to determine the IP address of the client
- **title:** Title of the application

### Theme

<p align="justify">
Easy Gate theme can be configured by providing colors for background and foreground. Theme changes will be applied immediately.
</p>

Example of a dark mode theme:

#### JSON

```json
"theme": {
  "background": "#1d1d1d",
  "foreground": "#ffffff",
}
```

#### YAML

```yml
theme:
  background: "#FFFFFF"
  foreground: "#000000"
```

### Groups

<p align="justify">
Group entries are used to define which users can see which items, by providing the user subnet. Group functionality is useful when dealing with both internal network and VPN users.

</p>

#### JSON

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

#### YAML

```yml
groups:
  - name: internal
    subnet: 192.168.1.1/24
  - name: vpn
    subnet: 10.8.1.1/24
```

### Services

<p align="justify">
A service entry is used to define a service that is available in the infrastructure. Each service has the following configurable parameters:

- **name:** the name of the service (ex. Internal Git, Jenkins, ...)
- **url:** the service url (must be a valid url starting with http(s)://)
- **category (optional):** service category (each service will appear under its own category).
- **groups:** list of groups associated to this service (defined in the groups section). If no group is provided the item can be seen by all users:
- **icon (optional):** the icon parameter accepts image URLs or data URI. If the icon parameter is not provided or empty, Easy Gate will try to fetch the service favicon and display it or fallback to a default icon.

</p>

#### JSON

```json
{
  "name": "Git",
  "url": "https://git.example.internal",
  "groups": [
      "vpn"
  ]
},
{
  "name": "Portainer",
  "url": "https://portainer.example.all",
  "category": "Test",
  "icon": "data:image/png;base64,[...]",
  "groups": []
}
```

#### YAML

```yml
- name: Git
  url: https://git.example.internal
  groups:
    - vpn
- name: Portainer
  url: https://portainer.example.all
  category: "Test"
  icon: data:image/png;base64,[...]
  groups: []
```

### Notes

<p align="justify">
A note entry is used to define a simple text note which has a title and a content. Each note has a name, the note content (text) and the groups that can see it (defined in the groups section). If no group is provided the item can be seen by all users:
</p>

#### JSON

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

#### YAML

```yml
- name: Simple note
  text: This is a simple note for vpn users
  groups:
    - vpn
- name: Global note
  text: This note will be visible to everyone
  groups: []
```

### Environment Variables

- **EASY_GATE_CONFIG_PATH:** Easy Gate configuration file path can be provided by this environment variable. The value will have precedence over the configuration file path provided in the command line.
- **EASY_GATE_CONFIG:** Insted of providing a configuration file, it is possible to provide the entire configuration as a JSON or YAML string in this environment variable. The content of this variable will have precedence over the configuration file.
