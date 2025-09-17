let fileInput = document.getElementById("file");

fileInput.addEventListener("change", () => {
  let file = fileInput.files[0];
  upload(file);
});

async function upload(file) {
  let formData = new FormData();
  formData.append("file", file);

  try {
    await fetch("http://localhost:8000/upload", {
      method: "POST",
      body: formData,
    });
  } catch (e) {
    console.log(e);
  }
}
