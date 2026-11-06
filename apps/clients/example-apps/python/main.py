import sys, os
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '../../clients/python')))
import openapi_client

def run():
    with openapi_client.ApiClient() as api_client:
        api = openapi_client.AdviceApi(api_client)
        try:
            api.v1_text_advice_get()
        except Exception as e:
            if hasattr(e, 'status') and e.status == 401:
                print("✅ Python Client E2E OK (Expected 401)")
                sys.exit(0)
            print(f"❌ Failed: {e}")
            sys.exit(1)

if __name__ == '__main__':
    run()
