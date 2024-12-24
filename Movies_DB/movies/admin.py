from django.contrib import admin
from .models import Collection, Movie  # Import your models here

# Optionally, create a custom admin class for the Collection model
class CollectionAdmin(admin.ModelAdmin):
    list_display = ('id', 'name')  # Specify which fields to display in the list view
    search_fields = ('name',)       # Add search functionality
    ordering = ('id',)               # Order by ID by default

class MovieAdmin(admin.ModelAdmin):
    list_display = ('id', 'external_id', 'title', 'genre',)  # Adjust to your actual fields
    search_fields = ('id','title',)                     # Add search functionality
    ordering = ('id',)                                   # Order by ID by default

# Register the models with the admin site
admin.site.register(Collection, CollectionAdmin)
admin.site.register(Movie, MovieAdmin)  # Corrected this line
