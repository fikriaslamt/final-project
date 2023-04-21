# final project - My Gram

simple social media REST Api \
host : ``` https://localhost ```

## endpoint 

**user** :
- register (POST) : ```/users/register``` \
  use json format  
  ```json
    {
        "username"  : "",
        "email"     : "",
        "password"  : "",
        "age"       : ""
    }
  ```
- login (POST) : ```/users/login``` \
  use json format  
  ```json
    {
        "email"     : "",
        "password"  : ""
    }
  ```

**social media** : 
- getAllSocialMedia (GET) : ``` /social-media  ```
- getOneSocialMedia (GET) : ``` /social-media/:id  ```
- createSocialMedia (POST) : ``` /social-media/create ``` \
use json format  
  ```json
    {
        "name"              : "",
        "social_media_url"  : ""
    }
  ```

- updateSocialMedia (PUT) : ``` /social-media/update/:id  ```\
use json format  
  ```json
    {
        "name"              : "",
        "social_media_url"  : ""
    }
  ```
- deleteSocialMedia (DELETE) : ``` /social-media/delete/:id  ```

**photos** : 
- getAllPhotos (GET) : ``` /photos  ```
- getOnePhoto (GET) : ``` /photos/:id  ```
- createPhoto (POST) : ``` /photos/create ``` \
use json format  
  ```json
    {
        "title"       : "",
        "caption"     : "",
        "photo_url"   : ""
    }
  ```

- updatePhoto (PUT) : ``` /photos/update/:id  ```\
use json format  
  ```json
    {
        "title"       : "",
        "caption"     : "",
        "photo_url"   : ""
    }
  ```
- deletePhoto (DELETE) : ``` /photos/delete/:id  ```


**comments** : 
- getAllComments (GET) : ``` /comments/:photo_id  ```
- getOneComment (GET) : ``` /comments/:id  ```
- createComment (POST) : ``` /comments/create/:photo_id ``` \
use json format  
  ```json
    {
        "message": "",
    }
  ```

- updateComment (PUT) : ``` /comments/update/:id  ```\
use json format  
  ```json
    {
        "message": "",
    }
  ```
- deleteComment (DELETE) : ``` /comments/delete/:id  ```





  
 
