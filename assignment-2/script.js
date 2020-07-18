var nameValue;
var fileValue;
var image;
var strUser;
c=()=>{
	nameValue=document.getElementById("nameUser").value;
	console.log(nameValue)
}
d=()=>{
	fileValue=document.getElementById("image").value;
	console.log(fileValue)
}
var loadFile = function(event) {
    image = document.getElementById('output');
    image.src = URL.createObjectURL(event.target.files[0]);
}; 
function getBase64Image(img) {
    var canvas = document.createElement("canvas");
    canvas.width = img.width;
    canvas.height = img.height;

    var ctx = canvas.getContext("2d");
    ctx.drawImage(img, 0, 0);

    var dataURL = canvas.toDataURL("image/png");

    return dataURL.replace(/^data:image\/(png|jpg);base64,/, "");
}
var imgData
hush=()=>{
	var bannerImage = document.getElementById('bannerImg');
	imgData = getBase64Image(bannerImage);
	localStorage.setItem("imgData", imgData);
}
window.onload=function(){
	var name=document.getElementById("nameUser");
	name.addEventListener("input", c);
}



