import java.util.HashMap;
import java.util.Map;
import java.util.Scanner;
import java.util.regex.Pattern;

class Main {
    public static final Pattern validInputPattern = Pattern.compile("^[IVXLCDM0-9]+ [\\+\\-\\*/] [IVXLCDM0-9]+$");
    public static final Map<Character, Integer> romanNumerals = new HashMap<>() {
        {
            put('I', 1);
            put('V', 5);
            put('X', 10);
            put('L', 50);
            put('C', 100);
            put('D', 500);
            put('M', 1000);
        }
    };

    public static String calc(String input) {
        try {
            if (!validInputPattern.matcher(input).matches()) {
                throw new IllegalArgumentException("Неверный формат выражения!");
            }

            String[] parts = input.split(" ");
            String num1Str = parts[0];
            String operatorStr = parts[1];
            String num2Str = parts[2];

            if (isRoman(num1Str) != isRoman(num2Str)) {
                throw new IllegalArgumentException("Использование разных систем счисления запрещено!");
            }

            int num1 = isRoman(num1Str) ? RomanConverter.toArabic(num1Str) : Integer.parseInt(num1Str);
            int num2 = isRoman(num2Str) ? RomanConverter.toArabic(num2Str) : Integer.parseInt(num2Str);

            int result;
            switch (operatorStr) {
                case "+":
                    result = num1 + num2;
                    break;
                case "-":
                    result = num1 - num2;
                    break;
                case "*":
                    result = num1 * num2;
                    break;
                case "/":
                    if (num2 == 0) {
                        throw new ArithmeticException("Деление на ноль!");
                    }
                    result = num1 / num2;
                    break;
                default:
                    throw new IllegalArgumentException("Неподдерживаемая операция!");
            }

            return isRoman(num1Str) ? RomanConverter.toRoman(result) : String.valueOf(result);
        } catch (NumberFormatException e) {
            throw new IllegalArgumentException("Введите корректные числа!");
        }
    }

    public static boolean isRoman(String str) {
        return str.matches("^[IVXLCDM]+$");
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        System.out.println("Добро пожаловать в Калькулятор!");
        System.out.println("Введите арифметическое выражение в формате: число оператор число");
        System.out.println("Для выхода введите 'exit'");

        while (true) {
            System.out.print("Выражение: ");
            String input = scanner.nextLine();

            if (input.equalsIgnoreCase("exit")) {
                System.out.println("До свидания!");
                break;
            }

            try {
                String result = calc(input);
                System.out.println("Результат: " + result);
            } catch (Exception e) {
                System.out.println("Ошибка: " + e.getMessage());
            }
        }

        scanner.close();
    }
}

class RomanConverter {
    public static String toRoman(int number) {
        if (number <= 0 || number >= 4000) {
            throw new IllegalArgumentException("Неверное значение для конвертации в римские числа.");
        }

        int[] values = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
        String[] numerals = {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};

        StringBuilder roman = new StringBuilder();

        int i = 0;
        while (number > 0) {
            if (number - values[i] >= 0) {
                roman.append(numerals[i]);
                number -= values[i];
            } else {
                i++;
            }
        }

        return roman.toString();
    }

    public static int toArabic(String roman) {
        int result = 0;
        int prevValue = 0;

        for (int i = roman.length() - 1; i >= 0; i--) {
            int value = Main.romanNumerals.getOrDefault(roman.charAt(i), 0);

            if (value < prevValue) {
                result -= value;
            } else {
                result += value;
            }

            prevValue = value;
        }

        return result;
    }
}

