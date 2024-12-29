package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Question struct {
    Question string
    Options  []string
    Answer   string
}

func main() {
    questions := []Question{
        {
            Question: "Какой автомобиль известен своим логотипом с четырьмя кольцами?",
            Options:  []string{"A) BMW", "B) Audi", "C) Mercedes", "D) Volkswagen"},
            Answer:   "B",
        },
        {
            Question: "Какой бренд производит автомобили под маркой 'Civic'?",
            Options:  []string{"A) Toyota", "B) Honda", "C) Nissan", "D) Ford"},
            Answer:   "B",
        },
        {
            Question: "Какой автомобиль называется 'Жигули'?",
            Options:  []string{"A) Lada", "B) Moskvitch", "C) GAZ", "D) Volga"},
            Answer:   "A",
        },
        {
            Question: "Какая марка выпускает модели 'Model S' и 'Model 3'?",
            Options:  []string{"A) BMW", "B) Tesla", "C) Ford", "D) Chevrolet"},
            Answer:   "B",
        },
        {
            Question: "Какой автомобиль считается первым серийным автомобилем в мире?",
            Options:  []string{"A) Ford Model T", "B) Benz Patent-Motorwagen", "C) Rolls-Royce Silver Ghost", "D) Vauxhall 25-hp"},
            Answer:   "B",
        },
    }

    score := 0
    reader := bufio.NewReader(os.Stdin)

    for _, question := range questions {
        fmt.Println(question.Question)
        for _, option := range question.Options {
            fmt.Println(option)
        }

        fmt.Print("Ваш ответ: ")
        answer, _ := reader.ReadString('\n')
        answer = strings.TrimSpace(answer)

        if strings.ToUpper(answer) == question.Answer {
            fmt.Println("Верно!")
            score++
        } else {
            fmt.Printf("Неверно! Правильный ответ: %s.\n", question.Answer)
        }
        fmt.Println()
    }

    fmt.Printf("Викторина окончена! Ваш счет: %d из %d\n", score, len(questions))
}
