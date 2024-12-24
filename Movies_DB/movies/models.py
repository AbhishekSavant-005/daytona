from django.db import models
from django.contrib.auth.models import User

class Movie(models.Model):
    external_id = models.IntegerField(unique=True)  # Unique identifier for the movie
    title = models.CharField(max_length=255)  # Title of the movie
    description = models.TextField()  # Description of the movie
    genre = models.CharField(max_length=100)  # Genre of the movie

    def __str__(self):
        return self.title  # String representation of the movie

class Collection(models.Model):
    name = models.CharField(max_length=255)  # Field for the collection name
    user = models.ForeignKey(User, on_delete=models.CASCADE, null=True, blank=True)  # User who owns the collection
    movies = models.ManyToManyField(Movie, related_name='collections')  # Many-to-many relationship with Movie

    def __str__(self):
        return self.name  # String representation of the collection
