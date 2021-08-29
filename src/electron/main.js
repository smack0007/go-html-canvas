const { app, BrowserWindow } = require("electron");
const fs = require("fs");
const path = require("path");

const HTML_FILE = "index.html";
let window = null;

function createWindow() {
  window = new BrowserWindow({
    show: false,
    autoHideMenuBar: true,
    useContentSize: true,
    width: 1024,
    height: 768,
    resizable: false,
    webPreferences: {
      preload: path.join(__dirname, "preload.js"),
    },
  });

  window.loadFile(HTML_FILE);
  window.webContents.openDevTools({ mode: "detach" });

  window.once("ready-to-show", () => {
    window.show();
  });
}

app.whenReady().then(() => {
  createWindow();
});

app.on("window-all-closed", function () {
  if (process.platform !== "darwin") app.quit();
});

fs.watch(__dirname, (eventType, fileName) => {
  console.info(`watch: ${fileName} => ${eventType}`);

  if (window !== null) {
    window.loadFile(HTML_FILE);
  //     if (hotReloadTimeout !== null) {
  //         clearTimeout(hotReloadTimeout);
  //         hotReloadTimeout = null;
  //     }

  //     hotReloadTimeout = (setTimeout(() => {
  //         if (window !== null) {
  //             console.info(`Reloading...`);

  //             window.webContents.send(ElectronEvent.hotReload);

  //             window.loadFile(HTML_FILE);
  //             hotReloadTimeout = null;
  //         }
  //     }, 500));
  }
});