import java.util.Scanner;

public class StringCalculator {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        try {
            while (true) {
                System.out.print("Введите выражение: ");
                String input = scanner.nextLine();

                String[] parts = input.split(" ");
                String operand1 = (parts.length == 4) ? parts[0] + " " + parts[1] : parts[0];
                String operator = (parts.length == 4) ? parts[2] : parts[1];
                String operand2 = (parts.length == 4) ? parts[3] : parts[2];

                if (parts.length < 3) {
                    throw new IllegalArgumentException("Неправильное выражение!");
                }

                if (!operand1.startsWith("\"") || !operand1.endsWith("\"")) {
                    throw new IllegalArgumentException("Первый операнд должен быть строкой в кавычках!");
                }

                operand1 = operand1.substring(1, operand1.length() - 1);
                String result = "";

                switch (operator) {
                    case "+":
                        if (isNumeric(operand2)) {
                            int num = Integer.parseInt(operand2);
                            result = "\""+ (operand1 + num)+ "\"";
                        } else {
                            if (!operand2.startsWith("\"") || !operand2.endsWith("\"")) {
                                throw new IllegalArgumentException("Второй операнд должен быть строкой в кавычках!");
                            }
                            operand2 = operand2.substring(1, operand2.length() - 1);
                            result = "\""+ (operand1 + operand2) + "\"";
                        }
                        break;
                    case "-":
                        if (isNumeric(operand2)) {
                            throw new IllegalArgumentException("Нельзя вычитать число из строки!");
                        } else {
                            if (!operand2.startsWith("\"") || !operand2.endsWith("\"")) {
                                throw new IllegalArgumentException("Второй операнд должен быть строкой в кавычках!");
                            }
                            operand2 = operand2.substring(1, operand2.length() - 1);
                            result = "\""+ (operand1.replaceFirst(operand2, "") )+ "\"";
                        }
                        break;
                    case "*":
                        int n = Integer.parseInt(operand2);
                        if (n < 1 || n > 10) {
                            throw new IllegalArgumentException("Число n должно быть от 1 до 10 включительно!");
                        }
                        StringBuilder sb = new StringBuilder();
                        for (int i = 0; i < n; i++) {
                            sb.append(operand1);
                        }
                        result = "\"" + sb.toString()+ "\"";
                        break;
                    case "/":
                        int divisor = Integer.parseInt(operand2);
                        if (divisor < 1 || divisor > 10) {
                            throw new IllegalArgumentException("Делитель должен быть от 1 до 10 включительно!");
                        }
                        int segmentLength = operand1.length() / divisor;
                        if (segmentLength == 0) {
                            result = "\"\"";
                        } else {
                            result = "\"" + operand1.substring(0, segmentLength) + "\"";
                        }
                        break;
                    default:
                        throw new IllegalArgumentException("Неподдерживаемая операция: " + operator);
                }

                if (result.length() > 40) {
                    result = result.substring(0, 37) + "...";
                }

                System.out.println("Результат: " + result);
            }
        } catch (IllegalArgumentException e) {
            System.out.println("Ошибка: " + e.getMessage());
        } finally {
            scanner.close();
        }
    }

    
    private static boolean isNumeric(String str) {
        try {
            Integer.parseInt(str);
            return true;
        } catch (NumberFormatException e) {
            return false;
        }
    }
}
