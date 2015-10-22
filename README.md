GoChat v0.1
===========

Project structure
-------------------

```
config					configuration functions
controllers				request handlers
helpers					different helpers
models					models
	response			response models
```


Requirements
------------
1. Go compiler
2. MySQL


Services
--------
```
:81/registration
```
Registrtation service.

| Parameter | Type | Description |
|-----------|------|-------------|
| login     | string | User's login|
| password  | string | User's password |

```
:81/login
```
Authorization service

| Parameter | Type | Description |
|-----------|------|-------------|
| login     | string | User's login|
| password  | string | User's password |

```
:81/messaging
```
Long-poll for message receive

| Parameter | Type | Description |
|-----------|------|-------------|
| from     | int | User's id|
| to  | int | User's id |


## The MIT License

The MIT License (MIT)

Copyright (c) 2015 Sergey Zinchenko, DataLayer.ru

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.