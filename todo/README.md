# DBaas vs DIY

## DBaas Google Cloud SQL

### Pros

- Autoscaling
- Automated backups and recovery
- High-availability with dub-second downtime maintenance
- Monitoring
- Security
- Vector search
- Migration service

- Set up is easy with [gcloud or Console](https://cloud.google.com/sql/docs/mysql/create-manage-databases#console).

- Backups and recovery
  - Automated backups daily, custom location is possible. Google charges lower rate for backup storage.
  - Easy to restore from backup with gcloud or Console

- High-availability
  - Data redundancy: If a zone or instance becomes unavailable, the data will be available for client applications. 

- Maintenance
  - Everything is updated automatically. Can result in minimal downtime.

- Monitoring and logging
  - Predefined dashboards for Cloud SQL monitoring. Custom dashboards are also possible.
  - Cloud monitoring can be used to set up alerts.
  - SQL queries are logged and can be viewed.

- Pricing
  - No need to keep specialized database engineer teams to manage database, so would probably be cheaper than self managed. Depends on project.
  - A highly available production database could cost from [250€ to 5000€](https://cloud.google.com/sql/docs/pricing-examples), which would probably still cost less than dedicating engineers yourself.

### Cons

- Could be overkill for small personal projects and would cost more.
- No full access to the servers and configurations.
- Risk of vendor lock in.

## PersistentVolumeClaims

## Pros


- Control
  - More granular access to resources.
  - Can be configured with any images.

- Pricing
  - Running database can be cheaper.
  - More control allows for more optimizations.

- No risk of vendor lock in.
 

## Cons

- Set up
  - More complex than Google Cloud SQL just to get the basic database running.

- Backups and Recovery
  - Must design and implement backup and recovery system.

- Maintenance
  - Nothing out of the box. Updating database engine must be done manually or with tools.
  - Must stay aware of security updates

- Monitoring and logging
  - Again everything must be implemented

- Pricing
  - If an organization must dedicate employee to implement and manage the database, the salary could be higher than using DBaas