function addElement(text) {
    if (text != '') {
        var elem = $('<li><span></span><button class="button_rm">Удалить</button></li>');
        $('span', elem).text(text);
        $('#root ul').append(elem);
        $('.button_rm', elem).click(function () {
          $(this).parent().remove();
        });
    }
}

$(document).ready(function () {
    $('#root').append('<ul></ul>');
    $('#root').append('<input id="add_task_input"><button id="add_task" class="input_button">Добавить</button>');
    $('#add_task').click(function() {
        addElement($('#add_task_input').val());
        $(this).parent().find('#add_task_input').val('');
    });
    $('#add_task_input').keypress(function (e) {
        if (e.which == 13) {
            e.preventDefault();
            addElement($('#add_task_input').val());
            $(this).val('');
        }
    })
    $(document).on('keypress', function(e) {
    var targ = e.target.tagName.toLowerCase();
    if ( e.which === 8 && targ != 'input')
        $('#root ul li:last').remove();
    });
    addElement("Сделать задание #3 по web-программированию");
})
