// 计算器逻辑 - 徐明 2024011
var out = document.getElementById("out");
var first = "";
var opr = "";
var reset = false;

function press(v) {
  if (reset) {
    out.textContent = v;
    reset = false;
  } else {
    out.textContent = out.textContent === "0" ? v : out.textContent + v;
  }
}

function oper(o) {
  first = out.textContent;
  opr = o;
  reset = true;
}

function calc() {
  if (opr === "") return;
  var a = parseFloat(first);
  var b = parseFloat(out.textContent);
  var r = 0;
  if (opr === "+") r = a + b;
  if (opr === "-") r = a - b;
  if (opr === "*") r = a * b;
  if (opr === "/") r = a / b;  // 未处理除零，会显示 Infinity
  out.textContent = r;
  first = ""; opr = "";
}

function ac() {
  out.textContent = "0"; first = ""; opr = ""; reset = false;
}

function del() {
  if (reset) return;
  out.textContent = out.textContent.length > 1 ? out.textContent.slice(0, -1) : "0";
}

// 绑定事件
function bind(id, fn) { document.getElementById(id).addEventListener("click", fn); }
bind("ac", ac); bind("del", del);
bind("div", function () { oper("/"); }); bind("mul", function () { oper("*"); });
bind("sub", function () { oper("-"); }); bind("add", function () { oper("+"); });
bind("eq", calc); bind("dot", function () { press("."); });
for (var i = 0; i < 10; i++) {
  (function (n) {
    bind("n" + n, function () { press(String(n)); });
  })(i);
}
