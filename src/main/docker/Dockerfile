#base package 
FROM java:8

#Create mount point tmp
VOLUME /tmp


ADD /target/spring-mysql-docker-0.0.1-SNAPSHOT.jar demo.jar
RUN bash -c 'touch /demo.jar'
ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/demo.jar"]