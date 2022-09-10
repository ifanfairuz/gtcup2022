import { createRoot } from "react-dom";

export function renderToId(id, element) {
  createRoot(document.getElementById(id)).render(element);
}

export function renderToRoot(element) {
  renderToId("root", element);
}

export function getData() {
  return window.DATA ?? {};
}
