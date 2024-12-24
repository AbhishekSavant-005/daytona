from django.urls import path, include
from rest_framework.routers import DefaultRouter
from .views import register, CollectionViewSet, MovieViewSet, list_movies
from movies.views import create_collection, get_request_count , get_collection, update_collection, delete_collection

router = DefaultRouter()
router.register(r'collections', CollectionViewSet, basename='collection')
router.register(r'movies', MovieViewSet, basename='movie')

urlpatterns = [
    path('register/', register, name='register'),
    path('movies/', list_movies, name='list_movies'),  # Move this line before the router urls
    path('', include(router.urls)),
    path('collections/create/', create_collection, name='create_collection'),
    path('collections/<int:collection_id>/', get_collection, name='get_collection'),
    path('collections/update/<int:collection_id>/', update_collection, name='update_collection'),
    path('collections/delete/<int:collection_id>/', delete_collection, name='delete_collection'),
    path('request_count/', get_request_count, name='get_request_count'),
]
