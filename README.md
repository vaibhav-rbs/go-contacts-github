# go-contacts-github





Setup POSTMAN client with correct token


  1. Set Environment:
    a. go to top right corner and click on gear icons and then select manage environment,
    b. now it should display you list of any existing environment if you have any,
    c. click on Add and set as following

      KEY     VALUE
      url :   https://gocontacts-vector.herokuapp.com/
      token:  <empty> this will be filled up later.
      d. it's advisable to save environment config 
  
  SET POST request url as {{url}}api/user/new which is equivalent to https://gocontacts-vector.herokuapp.com/api/user/new
  1. set Authorization to  No Auth 
  2. set as Headers:
      Authorization: Bearer {{token}}
      Content-Type: application/json
      
  3. enter some unique payload `{"email":"example@xyz.com","password":"abcd1234"}`
  4. now we will set up {{token}} automatically with new user creation or a sucessful login as follows
    ```
    var data = JSON.parse(responseBody);
    postman.clearGlobalVariable("token");
    postman.setGlobalVariable("token", data.account.token);
    ```
  5. this should have you and running
  
  

endpoints: Allowed methods
list users: https://gocontacts-vector.herokuapp.com/api/users(GET)
login: https://gocontacts-vector.herokuapp.com/api/user/login(POST)
register new user: https://gocontacts-vector.herokuapp.com/api/user/new(POST)



  
  var data = JSON.parse(responseBody);
  postman.clearGlobalVariable("token");
  postman.setGlobalVariable("token", data.account.token);
