const inputText = document.getElementById("input-text");
const buttGenarateQrcode = document.getElementById("btn-genarate-qrcode");

//http://localhost:8080
const URL_BASE = "/qrcode";


async function genarateQRCode() {
    const response = await fetch(URL_BASE, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            text: inputText.value
        })
    })

    const blobData = await response.blob();

    const fileReader = new FileReader();

    fileReader.onload = () => {
        const base64 = fileReader.result;
        document.getElementById("qrcode-image").src = base64;
        document.getElementById("qrcode-image").style.display = "block";
    }

    fileReader.readAsDataURL(blobData);
}

buttGenarateQrcode.addEventListener("click", genarateQRCode);