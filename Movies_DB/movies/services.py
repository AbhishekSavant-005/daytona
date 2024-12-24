# services.py
import requests
from django.conf import settings

import time

def fetch_movies_from_api():
    for attempt in range(3):  # Try 3 times
        try:
            response = requests.get(settings.MOVIE_API_URL, auth=(settings.API_USERNAME, settings.API_PASSWORD), timeout=10, verify=False)

            if response.status_code == 200:
                return response.json().get('results', [])
            else:
                print(f"Failed to fetch movies from API. Status code: {response.status_code}")
        except Exception as e:
            print(f"Error occurred while fetching movies: {e}")
        time.sleep(2)  # Wait before retrying
    return None
