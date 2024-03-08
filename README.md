# ConGo

Проект для изучения Go

<a href="https://highloadcup.ru/media/condition/accounts_rules_201218_4.html">Highload Cup 2018 (Task)</a>



Задача:
<p>В альтернативной реальности человечество решило создать и запустить глобальную систему по поиску "вторых половинок". Такая система призвана уменьшить количество одиноких людей в мире и способствовать созданию крепких семей.</p>

<p>Предлагается выступить в роли инженера, которому заказали создание прототипа подобной системы. Прототип должен как можно быстрее выдавать правильные ответы на запросы сторонних сервисов, которые делают что-то с ответами (например, отображают пользователям в красивых интерфейсах). По сути, он должен служить для внешних гипотетических сервисов функциональным API.</p>

# Описание необходимого API

API - это схемы http-запросов, которые должен обслуживать разработанный участником сервер. URL-ы строятся в соответствии с парадигмой REST. В угловых скобках указаны части URL, которые могут и будут меняться от запроса к запросу.

Во всех ответах от сервера учитываются заголовки Content-Type, Content-Length, Connection.

# Запросы выборки данных (GET):
<p>1. Получение списка пользователей: <strong>/accounts/filter/</strong></p>

<p>Данный метод API планируется использовать для поиска пользователей по заранее известным или желаемым полям. К примеру, кому-то захотелось посмотреть всех людей определённого возраста и пола, кто живёт в определённом городе.</p>

<p>В теле ответа ожидается структура <code>{"accounts": [ ... ]}</code> с пользователями, данные которых соответствуют указанным в GET-параметрах ограничениям. Для каждой подошедшей записи аккаунта не нужно передавать все известные о ней данные, а только поля id, email и те, что были использованы в запросе.</p>

<p>Пользователи в результате должны быть отсортированы по убыванию значений в поле id. Количество выбираемых записей ограничено обязательным GET-параметром limit.</p>

<p>Остальные GET-параметры формируются как <code>&lt;поле&gt;_&lt;предикат&gt;</code>. У разных полей могут использоваться только определённые фильтрующие предикаты, которые перечислены в таблице ниже. В данном запросе действие нескольких параметров складывается, то есть сначала фильтрация по одному, затем фильтрация результата по второму и т. д.</p>
<table class="table table-striped">
<tbody><tr><th>#</th><th>Название поля</th><th>Возможные предикаты с расшифровкой</th></tr>
<tr><td>1</td><td>sex</td><td><code>eq</code> - соответствие конкретному полу - "m" или "f";</td></tr>
<tr><td>2</td><td>email</td><td><code>domain</code> - выбрать всех, чьи email-ы имеют указанный домен;<br><code>lt</code> - выбрать всех, чьи email-ы лексикографически раньше;<br><code>gt</code> - то же, но лексикографически позже;</td></tr>
<tr><td>3</td><td>status</td><td><code>eq</code> - соответствие конкретному статусу;<br><code>neq</code> - выбрать всех, чей статус не равен указанному;</td></tr>
<tr><td>4</td><td>fname</td><td><code>eq</code> - соответствие конкретному имени;<br><code>any</code> - соответствие любому имени из перечисленных через запятую;<br><code>null</code> - выбрать всех, у кого указано имя (если 0) или не указано (если 1);</td></tr>
<tr><td>5</td><td>sname</td><td><code>eq</code> - соответствие конкретной фамилии;<br><code>starts</code> - выбрать всех, чьи фамилии начинаются с переданного префикса;<br><code>null</code> - выбрать всех, у кого указана фамилия (если 0) или не указана (если 1);</td></tr>
<tr><td>6</td><td>phone</td><td><code>code</code> - выбрать всех, у кого в телефоне конкретный код (три цифры в скобках);<br><code>null</code> - аналогично остальным полям;</td></tr>
<tr><td>7</td><td>country</td><td><code>eq</code> - всех, кто живёт в конкретной стране;<br><code>null</code> - аналогично;</td></tr>
<tr><td>8</td><td>city</td><td><code>eq</code> - всех, кто живёт в конкретном городе;<br><code>any</code> - в любом из перечисленных через запятую городов;<br><code>null</code> - аналогично;</td></tr>
<tr><td>9</td><td>birth</td><td><code>lt</code> - выбрать всех, кто родился до указанной даты;<br><code>gt</code> - после указанной даты;<br><code>year</code> - кто родился в указанном году;</td></tr>
<tr><td>10</td><td>interests</td><td><code>contains</code> - выбрать всех, у кого есть все перечисленные интересы;<br><code>any</code> - выбрать всех, у кого есть любой из перечисленных интересов;</td></tr>
<tr><td>11</td><td>likes</td><td><code>contains</code> - выбрать всех, кто лайкал всех перечисленных пользователей<br>&nbsp;(в значении - перечисленные через запятые id);</td></tr>
<tr><td>12</td><td>premium</td><td><code>now</code> - все у кого есть премиум на текущую дату;<br><code>null</code> - аналогично остальным;</td></tr>
</tbody></table>

Пример запроса и корректного ответа на него:
<pre><code class="hljs bash">GET: /accounts/filter/?status_neq=всё+сложно&amp;birth_lt=643972596&amp;country_eq=Индляндия&amp;<span class="hljs-built_in">limit</span>=5&amp;query_id=110</code></pre>
<pre><code class="json hljs">{
    <span class="hljs-attr">"accounts"</span>: [
        {
            <span class="hljs-attr">"email"</span>: <span class="hljs-string">"monnorakodehrenod@list.ru"</span>,
            <span class="hljs-attr">"country"</span>: <span class="hljs-string">"Индляндия"</span>,
            <span class="hljs-attr">"id"</span>: <span class="hljs-number">99270</span>,
            <span class="hljs-attr">"status"</span>: <span class="hljs-string">"заняты"</span>,
            <span class="hljs-attr">"birth"</span>: <span class="hljs-number">581863572</span>
        },{
            <span class="hljs-attr">"email"</span>: <span class="hljs-string">"erwirarhadmemeddifde@yahoo.com"</span>,
            <span class="hljs-attr">"country"</span>: <span class="hljs-string">"Индляндия"</span>,
            <span class="hljs-attr">"id"</span>: <span class="hljs-number">98881</span>,
            <span class="hljs-attr">"status"</span>: <span class="hljs-string">"свободны"</span>,
            <span class="hljs-attr">"birth"</span>: <span class="hljs-number">640015608</span>
        },{
            <span class="hljs-attr">"email"</span>: <span class="hljs-string">"rupewseor@rambler.ru"</span>,
            <span class="hljs-attr">"country"</span>: <span class="hljs-string">"Индляндия"</span>,
            <span class="hljs-attr">"id"</span>: <span class="hljs-number">98828</span>,
            <span class="hljs-attr">"status"</span>: <span class="hljs-string">"заняты"</span>,
            <span class="hljs-attr">"birth"</span>: <span class="hljs-number">604256977</span>
        },{
            <span class="hljs-attr">"email"</span>: <span class="hljs-string">"fiotnefaersohhev@inbox.ru"</span>,
            <span class="hljs-attr">"country"</span>: <span class="hljs-string">"Индляндия"</span>,
            <span class="hljs-attr">"id"</span>: <span class="hljs-number">98804</span>,
            <span class="hljs-attr">"status"</span>: <span class="hljs-string">"свободны"</span>,
            <span class="hljs-attr">"birth"</span>: <span class="hljs-number">596799123</span>
        },{
            <span class="hljs-attr">"email"</span>: <span class="hljs-string">"geslasereshedot@yahoo.com"</span>,
            <span class="hljs-attr">"country"</span>: <span class="hljs-string">"Индляндия"</span>,
            <span class="hljs-attr">"id"</span>: <span class="hljs-number">98718</span>,
            <span class="hljs-attr">"status"</span>: <span class="hljs-string">"свободны"</span>,
            <span class="hljs-attr">"birth"</span>: <span class="hljs-number">640919302</span>
        }
    ]
}</code></pre>

В случае неизвестного поля или неразрешённого предиката, в ответе ожидается код 400 с пустым телом. Во всех остальных случаях ожидается ответ 200, даже если ни одного пользователя не нашлось.
