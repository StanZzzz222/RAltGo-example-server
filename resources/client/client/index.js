import * as alt from "alt-client";

alt.onServer("hello", (name, age) => {
  alt.emitServer("hello", name, age)
});

alt.onServer("emitbenchmark", (...args) => {
  let users = JSON.parse(args[0])
  console.log(`Users Count - ${users.length}`)
});

alt.onServer("emitbenchmark:objects", (user) => {
  console.log(`Users Name - ${user.name}`)
});

alt.onServer("test_meta_data", () => {
  let player = alt.Player.local
  let testData = player.getSyncedMeta("test")
  let boolData = player.getSyncedMeta("bool")
  let playerData = player.getSyncedMeta("player")
  console.log(testData)
  console.log(boolData)
  console.log(playerData)
});
