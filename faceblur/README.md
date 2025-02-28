## How to run

1. Put .env files in place

2. Install what's needed

```bash
make install
```

3. Run the infra in a separate terminal

```bash
make run_local_infra
```

4. Create resources

```bash
make create_resources
make create_app_resources
```

5. Run apps, each command in a separate terminal

```bash
make run app=backend svc=api
```

```bash
make run app=backend svc=worker
```

```bash
make run app=dashboard
```
