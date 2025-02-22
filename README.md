# GoBoard

![GoBoard Logo](./assets/logo.png)
GoBoard is your local-first scheduler and to-do app, drawing inspiration from **Kanban boards** and **Neovim**. Built with **Raylib**, GoBoard is designed to be fast, lightweight, and keyboard-driven.

![GoBoard Demo](./assets/demo.gif) 

---

## Features

- **Local-First**: All data is stored locally on your machine.
- **Keyboard-Driven**: Designed for efficiency with minimal mouse usage.
- **Task Management**: Create, edit, and organize tasks with ease.
- **Inspired by Kanban and Neovim**: Combines the best of both worlds for productivity.

---

## Keybindings

| Key | Action                     |
|-----|----------------------------|
| `i` | Create a new task          |
| `\` | Save the current task      |
| `Esc` | Exit task creation mode |

---

## Task Format

When creating a new task, follow this format:

1. **First Line**: Task title.
2. **Second Line**: Task description.
3. **Third Line**: Start time (e.g., `2023-10-15 09:00`).
4. **Fourth Line**: End time (e.g., `2023-10-15 10:00`).
5. **Fifth Line**: Repeat days (using the `umtwrfs` format):
   - `u` = Sunday
   - `m` = Monday
   - `t` = Tuesday
   - `w` = Wednesday
   - `r` = Thursday
   - `f` = Friday
   - `s` = Saturday

Example:
Finish Project
Complete the final report
2023-10-15 09:00
2023-10-15 10:00
mtwr

## Under Development

GoBoard is still under active development. Contributions and feedback are welcome! Check out the [issues](https://github.com/yahiachahine/goboard/issues) to see what's being worked on.

---

## License

GoBoard is licensed under the [MIT License](./LICENSE).

---

## Acknowledgments

- Inspired by **Kanban boards** and **Neovim**.
- Built with **[Raylib](https://www.raylib.com/)**.

---

