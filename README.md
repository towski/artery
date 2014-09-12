asphalt
=======

The opposite of rails


=======
Decouple the controller, DB writer, and HTML templating. 

The actions never return html, only 302s or json. 

The html is served statically, but requests to actions can trigger html file rebuilds.

The DB writer lives in its own thread, so it never blocks the actions.
