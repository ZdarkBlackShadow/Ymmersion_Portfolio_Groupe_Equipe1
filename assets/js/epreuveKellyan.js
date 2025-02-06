document.addEventListener("DOMContentLoaded", function () {
    Prism.highlightAll();
  });
  document.addEventListener("DOMContentLoaded", () => {
    const columns = document.querySelectorAll(".column");

    columns.forEach((column, index) => {
      if (index % 2 === 0) {
        column.classList.add("slide-left");
      } else {
        column.classList.add("slide-right");
      }
    });
  });