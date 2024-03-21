function submitForm() {
    var form = document.getElementById("tagForm");
    form.action = window.location.href;
    form.submit();
}