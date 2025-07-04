

# 🎬 Kinopoisk Movie Search API Wrapper (Golang)

Простой прокси-сервер на Go, который оборачивает публичное API [kinopoisk.dev](https://kinopoisk.dev), позволяя искать фильмы по названию и получать краткую информацию с рейтингом и постером.

````markdown
---

## 🚀 Функциональность

- `GET /search?title=Название`  
  Ищет фильмы по названию, отображает:
  - Название
  - Год
  - ID
  - Рейтинг (KP)
  - Постер (HTML-картинка)

---

## 💻 Стек

- Go `net/http`
- JSON парсинг
- .env переменные (`joho/godotenv`)
- HTML-рендеринг на лету

---

## 📦 Установка

```bash
git clone https://github.com/yourname/kinopoiskAPI.git
cd kinopoiskAPI
go mod tidy
````

---

## 🔐 Настройка `.env`

Создай `.env` файл в корне проекта:

```env
KINOPOISK_API_KEY=your_api_key_here
```

API-ключ можно получить на [https://kinopoisk.dev](https://kinopoisk.dev).

---

## ▶️ Запуск

```bash
go run .
```

---

## 📡 Пример запроса

```http
GET http://localhost:3000/search?title=Брат
```

**Открой в браузере** — и увидишь красиво отрисованные заголовки с постерами и рейтингом.

---

## 💡 Идеи для улучшения

* Добавить `/movie?id=...` — подробный просмотр одного фильма
* Кэширование ответов
* HTML-шаблоны с `html/template`
* Swagger-документация
* Dockerfile + CI/CD

---

## 🧠 Автор

GrammSoli - Разработано в рамках pet-проекта с целью изучения работы с внешними API и Go в вебе.

---

## 📜 Лицензия

MIT
