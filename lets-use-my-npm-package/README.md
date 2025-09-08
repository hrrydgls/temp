# Using my package

Init npm like this in a empty folder:

```bash
npm init
```

Then a lot of enters to create the `pakcage.json` file.

Then set `"type": "module"` to make sure that this is a ESM.

And finally:

```js
import { greet } from "hrrydgls-greet"

console.log(greet("World"))
```
