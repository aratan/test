import pytest
from flask import Flask, url_for

@pytest.fixture
def client():
    app = Flask(__name__)
    app.config['TESTING'] = True

    with app.test_client() as client:
        yield client

def test_hello(client):
    response = client.get('/')
    assert response.status_code == 404
    
    
