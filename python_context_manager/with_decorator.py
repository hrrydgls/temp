from contextlib import contextmanager

@contextmanager
def my_context():
    # print("Entering")
    try:
        yield
    except Exception:
        raise
    finally:
        print("Exiting")

# with my_context():
#     print("Inside context")

with my_context():
    1 / 0