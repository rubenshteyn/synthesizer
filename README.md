# synthesizer

Эмулятор устройства для создания олигонуклеотидов. Проект написанный на Vue3. В качестве стейт-менеджера был использован Vuex. В качестве хранилища данных LocalStorage

## Stack
Vue3, Vuex, Vite

## Описание

У синтезатора есть 4 состояния: Занят, Бездействует, На обслуживании, Редактирование. Чтобы редактировать или удалить формулу в очереди задач, нажмите на нужную, синтезатор завершит текущую задачу и получит статус "Редактирование". По умолчанию при создании задачи её приоритет "Medium". Задачу, которая обрабатывается синтезатором(подсвечена зелёным) нельзя редактировать или удалять. Задачи выполняются по приоритету. На колбах видно который из нуклеотидов сейчас синтезируется в текущей формуле

# Screenshots
![Image alt](https://github.com/rubenshteyn/synthesizer/blob/main/src/assets/img/screenshots/frame.png)
![Image alt](https://github.com/rubenshteyn/synthesizer/blob/main/src/assets/img/screenshots/modal.png)
![Image alt](https://github.com/rubenshteyn/synthesizer/blob/main/src/assets/img/screenshots/changeForm.png)

## Project Setup

```sh
npm i
```

### Compile and Hot-Reload for Development

```sh
npm start
```

### Compile and Minify for Production

```sh
npm run build
```
