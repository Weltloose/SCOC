## 任务

模仿 Github,设计一个博客网站的 API

## 实现

### Authentication

```bash
curl -u "username" https://api.example.com 
```

### Failed login

Authenticating with invalid credentials will return `401 Unauthorized`:

```bash
curl -i https://api.example.com -u foo:bar
HTTP/1.1 401 Unauthorized
{
  "message": "Bad credentials",
  "documentation_url": "https://developer.example.com"
}
```

### Get All Articals User Owned

```bash
GET /user/{username}/articals
```

```json
Status：200 OK
--------------------------
{
  "total_count": 0,
  "items": [
    {
      "articleID": 0,
      "title": "",
      "article_url": "",
      "description": "",
      "private":"",
      "stars": 0,
      "forked": 0,
      "rank": 0,
      "createdAt": 0000-00-00 00:00:00,
      "updatedAt": 0000-00-00 00:00:00,
      "brief_content":"",
    },
    {
        ...
    },
  ]
}
```

### Get Public Artical with Given ArticalID

```bash
GET /articalID/{articalID}
```

```bash
curl -i "https://api.examle.com/api/articalID/{articalID}"
```

ON SUCCESS

```json
Status：200 OK
--------------------------
{
    "title": "",
    "content": "",
	"owner": {
        "username": "",
        "rank": 0,
        "articals": 0,
    }
    "createdAt": 0000-00-00 00:00:00,
    "updatedAt": 0000-00-00 00:00:00,
    "rank": 0,
    "stars": 0,
    "forked": 0,
}
```

ON FAILURE

```json
Status：401 Unauthorized
--------------------------
{
    "message": "Not Authed. ",
    "documentation_url": "https://developer.example.com",
}
```

```json
Status：404 Not Found
--------------------------
{
    "message": "Not Found. ",
    "documentation_url": "https://developer.example.com",
}
```

### Modified Certain Artical

```bash
PATCH /username/{username}/passwd/{passwd}/articalID
```

```bash
curl -i  -d '{"content":""}' "https://api.examle.com/api//username/{username}/passwd/{passwd}/articalID/{articalID}"
```

ON SUCCESS

```json
Status：200 OK
--------------------------
{
    "message": "MODIFIED SUCCESS!",
}
```

ON FAILURE

```json
Status：401 Unauthorized
--------------------------
{
    "message": "Not Authed. ",
    "documentation_url": "https://developer.example.com",
}
```

```json
Status：404 Not Found
--------------------------
{
    "message": "Not Found. ",
    "documentation_url": "https://developer.example.com",
}
```

### Delete Certain Artical

```bash
DELETE /username/{username}/passwd/{passwd}/articalID
```

```bash
curl -i "https://api.examle.com/api/username/{username}/passwd/{passwd}/articalID/{articalID}"
```

ON SUCCESS

```json
Status：200 OK
--------------------------
{
    "message": "DELETE SUCCESS!",
}
```

ON FAILURE

```json
Status：401 Unauthorized
--------------------------
{
    "message": "Not Authed. ",
    "documentation_url": "https://developer.example.com",
}
```

```json
Status：404 Not Found
--------------------------
{
    "message": "Not Found. ",
    "documentation_url": "https://developer.example.com",
}
```

  ### Create Artical

```bash
POST /username/{username}/passwd/{passwd}/articalID/{articalID}
```

```bash
curl -i -d '{"title": "", "content": "", }' '"https://api.examle.com/api/username/{username}/passwd/{passwd}/articalID/{articalID}"
```

ON SUCCESS

```json
Status：200 OK
--------------------------
{
    "message": "Create Success!",
    "createdAt": 0000-00-00 00:00:00,
    "updatedAt": 0000-00-00 00:00:00,
}
```

ON FAILURE

```json
Status：401 Unauthorized
--------------------------
{
    "message": "Not Authed. ",
    "documentation_url": "https://developer.example.com",
}
```

### Star an Artical

```bash
PATCH /articalID/{articalID}
```

```bash
curl -i "https://api.examle.com/api/articalID/{articalID}"
```

```json
Status：200 OK
--------------------------
{
    "message": "Star Success!",
}
```

### Fork an Artical

```bash
POST /username/{username}/passwd/{passwd}/articalID/{articalID}/fork
```

```bash
curl -i -d '{"articalID": 0,}' "https://api.examle.com/api/username/{username}/passwd/{passwd}/articalID/{articalID}/fork"
```

ON SUCCESS

```json
Status：200 OK
--------------------------
{
    "message": "FORK SUCCESS!"
	"forkFrom": {
        "username": "",
        "rank": 0,
        "articals": 0,
    }
    "createdAt": 0000-00-00 00:00:00,
    "updatedAt": 0000-00-00 00:00:00,
}
```

ON FAILURE

```json
Status：401 Unauthorized
--------------------------
{
    "message": "Not Authed. ",
    "documentation_url": "https://developer.example.com",
}
```

```json
Status：404 Not Found
--------------------------
{
    "message": "Not Found. ",
    "documentation_url": "https://developer.example.com",
}
```

### 

### 