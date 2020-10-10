# microBlog

A simple Go application that as of now serves simple html pages using go's html/template package.
All packages are native to go. 

Requirment
  * I have built this on go 1.15.2 
  
Notes
  * It runs on port 8080 by default
  * Make sure to keep the directory structure data, static folders must be in the same directory as the application
  * To add a new page you have to  add it to data/mapping.json
  
  eg : 
      {
		      "article_title": "apple-a-day",
		      "html_file": "apple.html",
		      "layout_file": "apple_layout.html",
		      "article_name": "An Apple a day"
	    }
      
  * Have to restart the application to load any configuration changes in data/mapping.json
 
Build & Run

  * in the root directory "go build -o microblog"
  * run "./microblog" to  start the application
 
