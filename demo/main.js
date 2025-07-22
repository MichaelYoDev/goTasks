const output = document.getElementById("output");
const input = document.getElementById("commandInput");
let tasks = [];

input.addEventListener("keydown", (e) => {
  if (e.key === "Enter") {
    const command = input.value.trim();
    input.value = "";
    handleCommand(command);
  }
});

function handleCommand(cmd) {
  writeLine(`$ ${cmd}`);

  if (cmd.startsWith("goTasks add ")) {
    const task = cmd.replace("goTasks add ", "").replace(/"/g, "");
    tasks.push({ description: task, done: false });
    writeLine(`  Task added: ${task}`);
  } else if (cmd === "goTasks list") {
    const list = tasks.filter((t) => !t.done);
    if (list.length === 0) {
      writeLine("  No tasks.");
    } else {
      list.forEach((t, i) => writeLine(`  ${i + 1}. ${t.description}`));
    }
  } else if (cmd === "goTasks list -a" || cmd === "goTasks --all") {
    if (tasks.length === 0) {
      writeLine("  No tasks.");
    } else {
      tasks.forEach((t, i) => {
        const status = t.done ? "✓" : "✗";
        writeLine(`  ${i + 1}. ${t.description} [${status}]`);
      });
    }
  } else if (cmd.startsWith("goTasks complete ")) {
    const id = parseInt(cmd.split(" ")[2]) - 1;
    if (!isNaN(id) && tasks[id]) {
      tasks[id].done = true;
      writeLine(`  Task ${id + 1} marked as complete.`);
    } else {
      writeLine("  Invalid task ID.");
    }
  } else if (cmd.startsWith("goTasks delete ")) {
    const id = parseInt(cmd.split(" ")[2]) - 1;
    if (!isNaN(id) && tasks[id]) {
      tasks.splice(id, 1);
      writeLine(`  Task ${id + 1} deleted.`);
    } else {
      writeLine("  Invalid task ID.");
    }
  } else {
    writeLine("  Command not recognized.");
  }
}

function writeLine(text) {
  output.innerHTML += text + "\n";
  window.scrollTo(0, document.body.scrollHeight);
}
