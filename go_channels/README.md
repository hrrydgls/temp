# Go Channels

I have simulated a web page download and extract scenario. so `download` function will download and send output to the `extract` to do sth on top of output of the first action...

I get sth like this at the end as output:

```bash
git:(master) ✗ go build . && ./temp -command=channel
A content downloaded
B content downloaded
A content extraced
B content extraced
C content downloaded
C content extraced
```

As you see one goroutine can send the value to the second one as soon as it finishes! so they go ahead in parallel without requring to finish the first one to go to the next!
