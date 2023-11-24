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
                BufferedReader in = new BufferedReader(new InputStreamReader(connection.getInputStream()));
                StringBuilder response = new StringBuilder();
                String inputLine;

                while ((inputLine = in.readLine()) != null) {
                    response.append(inputLine);
                }
                in.close();
                String imageData = response.toString();
                System.out.println("AFTER IMAGE READING");
                byte[] imageBytes = imageData.getBytes(); // Assuming the response is already in binary format (byte array)
                System.out.println(Arrays.toString(imageBytes));
                BufferedImage img = ImageIO.read(new ByteArrayInputStream(imageBytes));
                File outputfile = new File("/images/image" + count + ".png");
                ImageIO.write(img, "png", outputfile);
            }
            else {
                System.out.println("HTTP GET request failed with response code: " + responseCode);
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