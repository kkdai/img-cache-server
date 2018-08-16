img-cache-server: A web image trasfer to HTTPS link
==============

 [![GoDoc](https://godoc.org/github.com/kkdai/img-cache-server?status.svg)](https://godoc.org/github.com/kkdai/img-cache-server)  [![Build Status](https://travis-ci.org/kkdai/img-cache-server.svg?branch=master)](https://travis-ci.org/kkdai/img-cache-server)



How to use it
=============

- "https://YOUR_ADDR/url? + `Your_Http_Img_Address`": Return a id of image cache.  (Img_Cache_ID)
- "https://YOUR_ADDR/img? + `Img_Cache_ID`.jpg": Cache image content.
- "https://YOUR_ADDR/go? + `Your_Http_Img_Address`": Provide https image on-fly, it may take a while if your image size is large.

So you can easily to cache any HTTP image and forward to HTTPS if you hosted on Heroko.

Enjoy it.

Installation and Usage
=============


[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)


Inspired by
---------------

- [imgproxy](https://evilmartians.com/chronicles/introducing-imgproxy)

License
---------------

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

