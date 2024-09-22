import * as alt from "alt-client";

alt.onServer("hello", (name, age) => {
  alt.emitServer("hello", name, age)
});
