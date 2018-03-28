import sys
import datetime
import json
import writer
import pprint

def int_or_none(v):
    if v is None or v == '':
        return None
    return int(v)

def float_or_none(v):
    if v is None or v == '':
        return None
    return float(v)

def date_or_none(v):
    if v is None or v == '':
        return None
    v = int(v)
    y = v / 10000
    m = (v % 10000) / 100
    d = v % 100
    return datetime.date(y, m, d)

def make_bool(v):
    return v != ''

def trim(obj):
    if obj is None or obj == '':
        return None
    if isinstance(obj, list):
        for i in reversed(xrange(len(obj))):
            obj[i] = trim(obj[i])
            if obj[i] is None and len(obj) == i+1:
                obj.pop(i)
        if len(obj) == 0:
            return None
        return obj
    if isinstance(obj, dict):
        for k, v in obj.items():
            v = trim(v)
            if v is None:
                del(obj[k])
            else:
                obj[k] = v
        if len(obj) == 0:
            return None
        return obj
    return obj

def parse_hdr(record):
    data = {
        'ef_type': record[1],
        'fec_version': record[2],
        'software': {
            'name': record[3],
            'version': record[4],
        },
        'report_id': record[5],
        'report_number': record[6],
    }
    if len(record) > 7:
        data['comment'] = record[7]
    return data

def parse_f3x(record):
    data = {
        'form_type': record[0],
        'filer_id': record[1],
        'committee_name': record[2],
        'change_of_address': make_bool(record[3]),
        'address': parse_address(record[4:]),
        'report_code': record[9],
        'election': {
            'code': record[10],
            'date': date_or_none(record[11]),
            'state': record[12],
        },
        'period': {
            'start': date_or_none(record[13]),
            'end': date_or_none(record[14]),
        },
        'qualified_committee': make_bool(record[15]),
        'treasurer': parse_name(record[16:]),
        'date_signed': date_or_none(record[21]),
        'entries': {
            '6': {
                'a': {
                    'ytd': float_or_none(record[73]),
                    'year': int_or_none(record[74]),
                },
                'b': {
                    'period': float_or_none(record[22]),
                },
                'c': {
                    'period': float_or_none(record[23]),
                    'ytd': float_or_none(record[75]),
                },
                'd': {
                    'period': float_or_none(record[24]),
                    'ytd': float_or_none(record[76]),
                },
            },
            '7': {
                'period': float_or_none(record[25]),
                'ytd': float_or_none(record[77]),
            },
            '8': {
                'period': float_or_none(record[26]),
                'ytd': float_or_none(record[78]),
            },
            '9': {
                'period': float_or_none(record[27]),
            },
            '10': {
                'period': float_or_none(record[28]),
            },
            '11': {
                'a': {
                    'i': {
                        'period': float_or_none(record[29]),
                        'ytd': float_or_none(record[79]),
                    },
                    'ii': {
                        'period': float_or_none(record[30]),
                        'ytd': float_or_none(record[80]),
                    },
                    'iii': {
                        'period': float_or_none(record[31]),
                        'ytd': float_or_none(record[81]),
                    },
                },
                'b': {
                    'period': float_or_none(record[32]),
                    'ytd': float_or_none(record[82]),
                },
                'c': {
                    'period': float_or_none(record[33]),
                    'ytd': float_or_none(record[83]),
                },
                'd': {
                    'period': float_or_none(record[34]),
                    'ytd': float_or_none(record[84]),
                },
            },
            '12': {
                'period': float_or_none(record[35]),
                'ytd': float_or_none(record[85]),
            },
            '13': {
                'period': float_or_none(record[36]),
                'ytd': float_or_none(record[86]),
            },
            '14': {
                'period': float_or_none(record[37]),
                'ytd': float_or_none(record[87]),
            },
            '15': {
                'period': float_or_none(record[38]),
                'ytd': float_or_none(record[88]),
            },
            '16': {
                'period': float_or_none(record[39]),
                'ytd': float_or_none(record[89]),
            },
            '17': {
                'period': float_or_none(record[40]),
                'ytd': float_or_none(record[90]),
            },
            '18': {
                'a': {
                    'period': float_or_none(record[41]),
                    'ytd': float_or_none(record[91]),
                },
                'b': {
                    'period': float_or_none(record[42]),
                    'ytd': float_or_none(record[92]),
                },
                'c': {
                    'period': float_or_none(record[43]),
                    'ytd': float_or_none(record[93]),
                },
            },
            '19': {
                'period': float_or_none(record[44]),
                'ytd': float_or_none(record[94]),
            },
            '20': {
                'period': float_or_none(record[45]),
                'ytd': float_or_none(record[95]),
            },
            '21': {
                'a': {
                    'i': {
                        'period': float_or_none(record[46]),
                        'ytd': float_or_none(record[96]),
                    },
                    'ii': {
                        'period': float_or_none(record[47]),
                        'ytd': float_or_none(record[97]),
                    },
                },
                'b': {
                    'period': float_or_none(record[48]),
                    'ytd': float_or_none(record[98]),
                },
                'c': {
                    'period': float_or_none(record[49]),
                    'ytd': float_or_none(record[99]),
                },
            },
            '22': {
                'period': float_or_none(record[50]),
                'ytd': float_or_none(record[100]),
            },
            '23': {
                'period': float_or_none(record[51]),
                'ytd': float_or_none(record[101]),
            },
            '24': {
                'period': float_or_none(record[52]),
                'ytd': float_or_none(record[102]),
            },
            '25': {
                'period': float_or_none(record[53]),
                'ytd': float_or_none(record[103]),
            },
            '26': {
                'period': float_or_none(record[54]),
                'ytd': float_or_none(record[104]),
            },
            '27': {
                'period': float_or_none(record[55]),
                'ytd': float_or_none(record[105]),
            },
            '28': {
                'a': {
                    'period': float_or_none(record[56]),
                    'ytd': float_or_none(record[106]),
                },
                'b': {
                    'period': float_or_none(record[57]),
                    'ytd': float_or_none(record[107]),
                },
                'c': {
                    'period': float_or_none(record[58]),
                    'ytd': float_or_none(record[108]),
                },
                'd': {
                    'period': float_or_none(record[59]),
                    'ytd': float_or_none(record[109]),
                },
            },
            '29': {
                'period': float_or_none(record[60]),
                'ytd': float_or_none(record[110]),
            },
            '30': {
                'a': {
                    'i': {
                        'period': float_or_none(record[61]),
                        'ytd': float_or_none(record[111]),
                    },
                    'ii': {
                        'period': float_or_none(record[62]),
                        'ytd': float_or_none(record[112]),
                    },
                },
                'b': {
                    'period': float_or_none(record[63]),
                    'ytd': float_or_none(record[113]),
                },
                'c': {
                    'period': float_or_none(record[64]),
                    'ytd': float_or_none(record[114]),
                },
            },
            '31': {
                'period': float_or_none(record[65]),
                'ytd': float_or_none(record[115]),
            },
            '32': {
                'period': float_or_none(record[66]),
                'ytd': float_or_none(record[116]),
            },
            '33': {
                'period': float_or_none(record[67]),
                'ytd': float_or_none(record[117]),
            },
            '34': {
                'period': float_or_none(record[68]),
                'ytd': float_or_none(record[118]),
            },
            '35': {
                'period': float_or_none(record[69]),
                'ytd': float_or_none(record[119]),
            },
            '36': {
                'period': float_or_none(record[70]),
                'ytd': float_or_none(record[120]),
            },
            '37': {
                'period': float_or_none(record[71]),
                'ytd': float_or_none(record[121]),
            },
            '38': {
                'period': float_or_none(record[72]),
                'ytd': float_or_none(record[122]),
            },
        },
    }
    return data

def parse_name(record):
    return {
        'last': record[0],
        'first': record[1],
        'middle': record[2],
        'prefix': record[3],
        'suffix': record[4],
    }

def parse_address(record):
    return {
        'street': [record[0], record[1]],
        'city': record[2],
        'state': record[3],
        'zipcode': record[4],
    }

def parse_schedule_a(record):
    data = {
        'form_type': record[0],
        'filer_id': record[1],
        'transaction_id': record[2],
        'ref_transaction_id': record[3],
        'ref_schedule_name': record[4],
        'contributor': {
            'type': record[5],
            'organization': record[6],
            'name': parse_name(record[7:]),
            'address': parse_address(record[12:]),
            'employer': record[23],
            'occupation': record[24],
        },
        'election': {
            'code': record[17],
            'description': record[18],
        },
        'contribution': {
            'date': date_or_none(record[19]),
            'amount': float_or_none(record[20]),
            'aggregate': float_or_none(record[21]),
            'purpose': record[22],
        },
        'donor': {
            'committee': {
                'fec_id': record[25],
                'name': record[26],
            },
            'candidate': {
                'fec_id': record[27],
                'name': parse_name(record[28:]),
                'office': record[33],
                'state': record[34],
                'district': int_or_none(record[35]),
            },
            'conduit': {
                'name': record[36],
                'address': parse_address(record[37:]),
            },
        },
        'memo': {
            'code': record[42],
            'text': record[43],
        },
        'account_ref': record[44],
    }
    return data

def parse_schedule_b(record):
    data = {
        'form_type': record[0],
        'filer_id': record[1],
        'transaction_id': record[2],
        'ref_transaction_id': record[3],
        'ref_schedule_name': record[4],
        'payee': {
            'type': record[5],
            'organization': record[6],
            'name': parse_name(record[7:]),
            'address': parse_address(record[12:]),
        },
        'election': {
            'code': record[17],
            'description': record[18],
        },
        'expenditure': {
            'date': date_or_none(record[19]),
            'amount': float_or_none(record[20]),
            'refund': float_or_none(record[21]),
            'purpose': record[22],
            'category_code': record[23],
        },
        'beneficiary': {
            'committee': {
                'fec_id': record[24],
                'name': record[25],
            },
            'candidate': {
                'fec_id': record[26],
                'name': parse_name(record[27:]),
                'office': record[32],
                'state': record[33],
                'district': int_or_none(record[34]),
            },
            'conduit': {
                'name': record[35],
                'address': parse_address(record[36:]),
            },
        },
        'memo': {
            'code': record[41],
            'text': record[42],
        },
        'account_ref': record[43],
    }
    return data

def parse(fp):
    form = {}
    field_sep = chr(0x1c)
    for line in fp:
        record = line.strip().split(field_sep)
        if record[0] == 'HDR':
            form['header'] = parse_hdr(record)
        elif record[0].startswith('F3X'):
            form['f3x'] = parse_f3x(record)
        elif record[0].startswith('SA'):
            if 'schedule_a' not in form:
                form['schedule_a'] = []
            form['schedule_a'].append(parse_schedule_a(record))
        elif record[0].startswith('SB'):
            if 'schedule_b' not in form:
                form['schedule_b'] = []
            form['schedule_b'].append(parse_schedule_b(record))
        else:
            raise Exception('unknown record type %s' % record[0])
    return form

if '__main__' == __name__:
    form = trim(parse(sys.stdin))
    pprint.pprint(form)
    #fec = writer.format(form)
    #sys.stdout.write(fec)
    #json.dump(trim(parse(sys.stdin)), sys.stdout, indent=2)

