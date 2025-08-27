# Go File Utility 🚀

Простая утилита для работы с файлами на Go, аналог `wc`.

## 📦 Функции

- ✅ Подсчет строк (`-l`)
- ✅ Подсчет слов (`-w`) 
- ✅ Подсчет символов (`-m`)
- ✅ Подсчет байт (`-c`)

## 🛠 Установка

```bash
git clone https://github.com/ВАШ_USERNAME/go-fileutil.git
cd go-fileutil
go build -o fileutil.exe
```

## 🚀 Использование

```bash
# Показать все метрики
./fileutil.exe test.txt

# Только строки и слова
./fileutil.exe -l -w test.txt

# Только байты
./fileutil.exe -c test.txt

# Помощь
./fileutil.exe -h
```

## 📊 Пример вывода

```
$ ./fileutil.exe test.txt
Строк: 3
Слов: 7  
Символов: 47
Байт: 47
```

## 🏗 Структура проекта

```
go-fileutil/
├── .vscode/
│   └── launch.json    # Конфигурация VS Code
├── .gitignore         # Игнорируемые файлы
├── go.mod            # Модуль Go
├── main.go           # Основной код
├── test.txt          # Тестовый файл
└── README.md         # Этот файл
```

## 📝 Для чего этот проект?

Этот проект создан для изучения:
- Базового синтаксиса Go
- Работы с файловой системой
- Обработки аргументов командной строки
- Работы с Git и GitHub

## 👨‍💻 Автор

Korvin07

## 📄 Лицензия

MIT License