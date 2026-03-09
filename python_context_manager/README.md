# Context Manager

Sets up a context for a block and cleans up after automatically. 

Behind the scence uses it:

```python
with open("file.txt", "r") as f:
    data = f.read()
```