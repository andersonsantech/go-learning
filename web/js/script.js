$('#registration-form').submit(function(event) {
    event.preventDefault();

    var name = $('#name').val();
    var email = $('#email').val();
    var age = $('#age').val();

    var user = {
        name: name,
        email: email,
        age: age
    };

    $.ajax({
        url: 'http://localhost:8080/users',
        method: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(user),
        success: function() {
            alert('User registered successfully');
        },
        error: function() {
            alert('An error occurred while registering the user');
        }
    });
});

$(document).ready(function() {
    $.get("http://localhost:8080/users", function(data) {
        try {
            var jsonData = JSON.parse(data);
            if (Array.isArray(jsonData)) {
                jsonData.forEach(function(user) {
                    var row = $("<tr>").append(
                        $("<td>").text(user.id),
                        $("<td>").text(user.email)
                    );
                    $("#user-table-body").append(row);
                });
            } else {
                console.error("A resposta da API não é uma array.");
            }
        } catch (error) {
            console.error("Erro ao fazer o parsing da string para JSON:", error);
        }
    });
});