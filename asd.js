function asd(){
    var email
    var passwd
    email=document.getElementById("email").value
    passwd=document.getElementById("pass").value
    fetch("add?email="+email+"&pass="+passwd,).then(console.log("AXIOS MUKODIK"))
}