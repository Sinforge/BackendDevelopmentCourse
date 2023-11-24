import java.io.IOException;
import java.io.OutputStream;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
import java.net.URL;
import java.nio.charset.StandardCharsets;
import java.time.LocalDateTime;
import java.util.*;

public class Main {
    public static void main(String[] args) throws IOException, InterruptedException {
        String targetUrl = "http://container_name:8000/endpoint";
        while (true) {

            String requestBody = "{\"prices\":" + getRandomPricesArray().toString() + ",\"productType\":" + "\""
                    + ProductType.randomValue().toString().toLowerCase() + "\"}";
            System.out.println(requestBody);

            URL url = new URL(targetUrl);
            HttpURLConnection connection = (HttpURLConnection) url.openConnection();

            connection.setRequestMethod("POST");
            connection.setRequestProperty("Content-Type", "application/json");
            connection.setRequestProperty("Accept", "application/json");
            connection.setDoOutput(true);

            try (OutputStream os = connection.getOutputStream()){
                byte[] input = requestBody.getBytes(StandardCharsets.UTF_8);
                os.write(input, 0, input.length);
            }

            Thread.sleep(1000);
        }
    }

    private enum ProductType {
        WASHING_MACHINE,
        TELEPHONE,
        BOOK;

        private static final Random random = new Random();
        public static ProductType randomValue() {
            ProductType[] productTypes = ProductType.values();
            return productTypes[random.nextInt(productTypes.length)];
        }
    }

    private static List<Long> getRandomPricesArray() {
        Random random = new Random();
        int arrayLength = random.nextInt(50) + 5;
        List<Long> prices = new ArrayList<>(arrayLength);
        for (int i = 0; i < arrayLength; i++) {
            prices.add(random.nextLong(100000));
        }
        return prices;
    }
}