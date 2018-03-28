import sys
import datetime
import pprint

def format_value(v):
    if v is None:
        return ''
    if isinstance(v, bool):
        if v:
            return 'X'
        return ''
    if isinstance(v, datetime.date):
        return v.strftime('%Y%m%d')
    if isinstance(v, (int, long)):
        return '%d' % v
    if isinstance(v, float):
        return '%.2f' % v
    return v.lstrip().replace('"', '')
    return v

def format_row(row):
    if row is None or len(row) == 0:
        return ''
    return chr(0x1c).join(format_value(x) for x in row) + '\n'

def format_hdr(rec):
    data = [
        'HDR',
        rec.get('ef_type', 'FEC'),
        rec.get('fec_version', '8.2'),
        rec.get('software', {}).get('name', 'FoothillDemsFEC'),
        rec.get('software', {}).get('version', '0.1'),
        rec.get('report_id'),
        rec.get('report_number')
    ]
    if 'comment' in rec:
        data.append(rec.get('comment'))
    return data

def get_entry(entries, key):
    keys = key.split('.')
    final = keys.pop()
    for k in keys:
        entries = entries.get(k, {})
    return entries.get(final, 0.0)

def format_f3x(record):
    data = [
        record.get('form_type'),
        record.get('filer_id'),
        record.get('committee_name'),
        record.get('change_of_address'),
    ] + format_address(record.get('address', {})) + [
        record.get('report_code'),
        record.get('election', {}).get('code'),
        record.get('election', {}).get('date'),
        record.get('election', {}).get('state'),
        record.get('period', {}).get('start'),
        record.get('period', {}).get('end'),
        record.get('qualified_committee'),
    ] + format_name(record.get('treasurer', {})) + [
        record.get('date_signed'),
    ]
    entries = record.get('entries', {})
    period = [
        get_entry(entries, '6.b.period'),
        get_entry(entries, '6.c.period'),
        get_entry(entries, '6.d.period'),
        get_entry(entries, '7.period'),
        get_entry(entries, '8.period'),
        get_entry(entries, '9.period'),
        get_entry(entries, '10.period'),
        get_entry(entries, '11.a.i.period'),
        get_entry(entries, '11.a.ii.period'),
        get_entry(entries, '11.a.iii.period'),
        get_entry(entries, '11.b.period'),
        get_entry(entries, '11.c.period'),
        get_entry(entries, '11.d.period'),
        get_entry(entries, '12.period'),
        get_entry(entries, '13.period'),
        get_entry(entries, '14.period'),
        get_entry(entries, '15.period'),
        get_entry(entries, '16.period'),
        get_entry(entries, '17.period'),
        get_entry(entries, '18.a.period'),
        get_entry(entries, '18.b.period'),
        get_entry(entries, '18.c.period'),
        get_entry(entries, '19.period'),
        get_entry(entries, '20.period'),
        get_entry(entries, '21.a.i.period'),
        get_entry(entries, '21.a.ii.period'),
        get_entry(entries, '21.b.period'),
        get_entry(entries, '21.c.period'),
        get_entry(entries, '22.period'),
        get_entry(entries, '23.period'),
        get_entry(entries, '24.period'),
        get_entry(entries, '25.period'),
        get_entry(entries, '26.period'),
        get_entry(entries, '27.period'),
        get_entry(entries, '28.a.period'),
        get_entry(entries, '28.bperiod'),
        get_entry(entries, '28.c.period'),
        get_entry(entries, '28.d.period'),
        get_entry(entries, '29.period'),
        get_entry(entries, '30.a.i.period'),
        get_entry(entries, '30.a.ii.period'),
        get_entry(entries, '30.b.period'),
        get_entry(entries, '30.c.period'),
        get_entry(entries, '31.period'),
        get_entry(entries, '32.period'),
        get_entry(entries, '33.period'),
        get_entry(entries, '34.period'),
        get_entry(entries, '35.period'),
        get_entry(entries, '36.period'),
        get_entry(entries, '37.period'),
        get_entry(entries, '38.period'),
    ]
    ytd = [
        get_entry(entries, '6.a.ytd'),
        entries.get('6', {}).get('a', {}).get('year'),
        get_entry(entries, '6.c.ytd'),
        get_entry(entries, '6.d.ytd'),
        get_entry(entries, '7.ytd'),
        get_entry(entries, '8.ytd'),
        get_entry(entries, '11.a.i.ytd'),
        get_entry(entries, '11.a.ii.ytd'),
        get_entry(entries, '11.a.iii.ytd'),
        get_entry(entries, '11.b.ytd'),
        get_entry(entries, '11.c.ytd'),
        get_entry(entries, '11.d.ytd'),
        get_entry(entries, '12.ytd'),
        get_entry(entries, '13.ytd'),
        get_entry(entries, '14.ytd'),
        get_entry(entries, '15.ytd'),
        get_entry(entries, '16.ytd'),
        get_entry(entries, '17.ytd'),
        get_entry(entries, '18.a.ytd'),
        get_entry(entries, '18.b.ytd'),
        get_entry(entries, '18.c.ytd'),
        get_entry(entries, '19.ytd'),
        get_entry(entries, '20.ytd'),
        get_entry(entries, '21.a.i.ytd'),
        get_entry(entries, '21.a.ii.ytd'),
        get_entry(entries, '21.b.ytd'),
        get_entry(entries, '21.c.ytd'),
        get_entry(entries, '22.ytd'),
        get_entry(entries, '23.ytd'),
        get_entry(entries, '24.ytd'),
        get_entry(entries, '25.ytd'),
        get_entry(entries, '26.ytd'),
        get_entry(entries, '27.ytd'),
        get_entry(entries, '28.a.ytd'),
        get_entry(entries, '28.b.ytd'),
        get_entry(entries, '28.c.ytd'),
        get_entry(entries, '28.d.ytd'),
        get_entry(entries, '29.ytd'),
        get_entry(entries, '30.a.i.ytd'),
        get_entry(entries, '30.a.ii.ytd'),
        get_entry(entries, '30.b.ytd'),
        get_entry(entries, '30.c.ytd'),
        get_entry(entries, '31.ytd'),
        get_entry(entries, '32.ytd'),
        get_entry(entries, '33.ytd'),
        get_entry(entries, '34.ytd'),
        get_entry(entries, '35.ytd'),
        get_entry(entries, '36.ytd'),
        get_entry(entries, '37.ytd'),
        get_entry(entries, '38.ytd'),
    ]
    #pprint.pprint(data)
    #pprint.pprint(period)
    #pprint.pprint(ytd)
    return data + period + ytd

def format_name(name):
    return [
        name.get('last'),
        name.get('first'),
        name.get('middle'),
        name.get('prefix'),
        name.get('suffix'),
    ]

def format_address(addr):
    street = list(x for x in addr.get('street', []))
    while len(street) < 2:
        street.append(None)
    return [
        street[0],
        street[1],
        addr.get('city'),
        addr.get('state'),
        addr.get('zipcode'),
    ]

def format_contributor(rec):
    return [
        rec.get('type'),
        rec.get('organization'),
    ] + format_name(rec.get('name', {})) + format_address(rec.get('address', {}))

def format_contribution(rec):
    return [
        rec.get('date'),
        rec.get('amount'),
        rec.get('aggregate'),
        rec.get('purpose'),
    ]

def format_committee(rec):
    return [
        rec.get('fec_id'),
        rec.get('name'),
    ]

def format_candidate(rec):
    return [
        rec.get('fec_id'),
    ] + format_name(rec.get('name', {})) + [
        rec.get('office'),
        rec.get('state'),
        rec.get('district'),
    ]

def format_conduit(rec):
    return [
        rec.get('name'),
    ] + format_address(rec.get('address', {}))

def format_donor(rec):
    return format_committee(rec.get('committee', {})) + format_candidate(rec.get('candidate', {})) + format_conduit(rec.get('conduit', {}))

def format_memo(rec):
    return [
        rec.get('code'),
        rec.get('text'),
    ]

def format_schedule_a(record):
    return [
        record.get('form_type'),
        record.get('filer_id'),
        record.get('transaction_id'),
        record.get('ref_transaction_id'),
        record.get('ref_schedule_name'),
    ] + format_contributor(record.get('contributor', {})) + [
        record.get('election', {}).get('code'),
        record.get('election', {}).get('description'),
    ] + format_contribution(record.get('contribution', {})) + [
        record.get('contributor', {}).get('employer'),
        record.get('contributor', {}).get('occupation'),
    ] + format_donor(record.get('donor', {})) + format_memo(record.get('memo', {})) + [
        record.get('account_ref')
    ]

def format_expenditure(rec):
    return [
        rec.get('date'),
        rec.get('amount'),
        rec.get('refund'),
        rec.get('purpose'),
        rec.get('category_code'),
    ]

def format_schedule_b(record):
    return [
        record.get('form_type'),
        record.get('filer_id'),
        record.get('transaction_id'),
        record.get('ref_transaction_id'),
        record.get('ref_schedule_name'),
    ] + format_contributor(record.get('payee', {})) + [
        record.get('election', {}).get('code'),
        record.get('election', {}).get('description'),
    ] + format_expenditure(record.get('expenditure', {})) + format_donor(record.get('beneficiary', {})) + format_memo(record.get('memo', {})) + [
        record.get('account_ref')
    ]

def format(form):
    rows = [
        format_hdr(form.get('header', {})),
        format_f3x(form.get('f3x', {})),
    ] + list(format_schedule_a(x) for x in form.get('schedule_a', [])) + list(format_schedule_b(x) for x in form.get('schedule_b', []))
    return ''.join(format_row(row) for row in rows)

