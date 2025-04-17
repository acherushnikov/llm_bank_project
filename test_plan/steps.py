from behave import given, when, then

@given('Авторизоваться как клиент в ДБО')
def step_impl_auth_client(context):
    # TODO: implement login step
    pass

@when('Перейти к оформлению кредита')
def step_impl_open_credit(context):
    # TODO: implement credit application navigation
    pass

@when('Заполнить анкету и подтвердить заявку')
def step_impl_submit_application(context):
    # TODO: implement form filling and submission
    pass

@when('Проверить, что кредит оформлен и задан период охлаждения (отображается дата/время окончания)')
def step_impl_verify_cooling_period(context):
    # TODO: implement verification of cooling period
    pass

@when('Зайти в ДБО, открыть кредит')
@when('Открыть карточку нового кредита')
@when('Открыть детали по кредиту в ДБО')
def step_impl_open_credit_card(context):
    # TODO: implement opening credit card in DBO
    pass

@when('Нажать \'Отказаться от кредита\'')
def step_impl_click_refuse_credit(context):
    # TODO: implement refusal button click
    pass

@when('Подтвердить отказ (при необходимости — подписать ЭЦП)')
def step_impl_confirm_refusal(context):
    # TODO: implement EDS confirmation
    pass

@when('Перевести точную сумму основного долга на указанные реквизиты')
def step_impl_full_payment(context):
    # TODO: simulate payment of loan principal
    pass

@then('Кредит успешно оформлен, параметр периода охлаждения установлен и отображается')
@then('Отображается срок действия периода (дата и время окончания)')
@then('Заявление успешно подано, отображается статус \'Принято\'')
@then('Выводится сообщение: \'Период охлаждения истёк, отказ невозможен\'')
@then('Система отклоняет повторную попытку: \'Заявление уже подано\'')
@then('Кредит завершён, начисления отменены, статус договора: \'Аннулирован\'')
@then('Система уведомляет о неполном возврате, отказ не завершён')
@then('Всё сторнировано, остаток задолженности: 0 ₽')
@then('Заявление отображается в системе с каналом подачи: \'Офис\'')
@then('Запись в журнале есть: заявление, канал, дата, сумма')
@then('Отчёт содержит все проведённые действия по отказам')
def step_impl_generic_assertions(context):
    # TODO: implement expected result verification
    pass