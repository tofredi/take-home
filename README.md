# Take Home - Alfredo Martinez

This tool makes http requests and prints the address of the request along with the
MD5 hash of the response.

## Usage:

- Compile packages and dependencies

  ```
  $ go build
  ```

- Run tool

  ```
  $ ./myhttp [-parallel n] addr1 addr2 ... addrN
  ```

- Examples
  ```
  $ ./myhttp http://www.facebook.com http://google.com
  http://google.com 38578ccb3a5583e785d277717c54a59c
  http://www.facebook.com 3822f73f3a8536991f5ef2daaf80e7ce
  ```

  ```
  $ ./myhttp facebook.com
  http://facebook.com d06bd2a0382e4dc210f74bef5dc9b272
  ```

  ```
  $ ./myhttp -parallel 3 facebook.com google.com facebook.com yahoo.com yandex.com twitter.com \
  reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com
  http://google.com e508d48fdec3ef1d46ec4ef347f0cd5f
  http://facebook.com e6cbc3e0b7f33b92a2eceaf23e190852
  http://facebook.com 4c2fbdd82d1e9228d36b5d6c91a8b59f
  http://twitter.com 4dbbe6d0fb151cf514703ae1428278e6
  http://yahoo.com 916819a7360c7d8dd644141dea5e0cbd
  http://yandex.com fb4762c350f9ceb8437fa0bcdba13c1f
  http://reddit.com/r/funny 4971565f5afcd1eb5becbb8ffc9699b3
  http://baroquemusiclibrary.com 84aca90e1570bb33c0e27733a3dba5ae
  http://reddit.com/r/notfunny e14694538aa7483ea01c8d88cb1506f5
  
  ```
  

- Run unit tests
  ```
  $ go test
  PASS
  ok      myhttp  0.545s
  ```
