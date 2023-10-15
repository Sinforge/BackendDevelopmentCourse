FROM bellsoft/liberica-openjdk-alpine:18.0.2.1-1
WORKDIR /app
COPY build/libs/Auth-0.0.1-SNAPSHOT.jar /app/Auth-0.0.1-SNAPSHOT.jar
EXPOSE 8081
CMD ["java", "-jar", "Auth-0.0.1-SNAPSHOT.jar"]