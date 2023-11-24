import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.io.*;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
import java.net.URL;
import java.nio.charset.StandardCharsets;
import java.time.LocalDateTime;
import java.util.*;

public class Main {
     public static void main(String[] args) throws IOException, InterruptedException {
        String targetUrl = "http://analyticsservice:8000/greeting/analytics/diagram/";
        int count = 0;
        for (; ; count++) {
            String requestBody = "{\"prices\":" + getRandomPricesArray() + ",\"productType\":" + "\""
                    + ProductType.randomValue().toString().toLowerCase() + "\"}";
            System.out.println(requestBody);

            URL url = new URL(targetUrl);
            HttpURLConnection connection = (HttpURLConnection) url.openConnection();

            System.out.println("BEFORE REQUEST");
            connection.setRequestMethod("POST");
            connection.setRequestProperty("Content-Type", "application/json");
            connection.setRequestProperty("Accept", "image/png");
            connection.setDoOutput(true);
            connection.setDoInput(true);

            try (OutputStream os = connection.getOutputStream()) {
                byte[] input = requestBody.getBytes(StandardCharsets.UTF_8);
                os.write(input, 0, input.length);
            }
            System.out.println("DATA WRITTEN TO REQUEST");
            int responseCode = connection.getResponseCode();
            System.out.println("AFTER REQUEST, BEFORE RESPONSE READING");
            if (responseCode == HttpURLConnection.HTTP_OK) {
                try (InputStream is = connection.getInputStream()) {
                    BufferedImage img = ImageIO.read(is);

                    // Save the image to a file
                    File outputfile = new File("image" + count + ".png");
                    ImageIO.write(img, "png", outputfile);

                    System.out.println("Image saved successfully.");
                }
            } else {
                System.out.println("HTTP POST request failed with response code: " + responseCode);
            }
            connection.disconnect();
            Thread.sleep(30000);
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