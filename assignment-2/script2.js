var dataImage = localStorage.getItem('imgData');
bannerImg = document.getElementById('displayImage');
bannerImg.src = "data:image/png;base64," + dataImage;
