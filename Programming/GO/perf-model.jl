users = 10e6
messages_per_user_per_day = 20
transactions_per_message = 2
domains_per_transaction = 2

messages_per_day = users * messages_per_user_per_day

num_st_shards = 48

transactions_per_day = users * messages_per_user_per_day * transactions_per_message
transactions_per_hour = transactions_per_day / 24
transactions_per_minute = transactions_per_hour / 60
transactions_per_second = transactions_per_hour / 3600

domains_per_minute = transactions_per_minute * domains_per_transaction
domains_per_minute_per_st_shard = domains_per_minute / num_st_shards
domains_per_second_per_st_shard = domains_per_minute_per_st_shard / 60

st_flush_period_min = 30

domains_per_flush_period = domains_per_minute * st_flush_period_min
domains_per_flush_period_per_st_shard = domains_per_flush_period / num_st_shards

active_groups_per_user = 5
avg_group_size = 3
active_groups = users * active_groups_per_user / avg_group_size

num_domains = users + active_groups
avg_domain_size_bytes = 100 * 1024
database_size_gb = num_domains * avg_domain_size_bytes / 1024 / 1024 / 1024
database_size_per_shard_gb = database_size_gb / num_st_shards